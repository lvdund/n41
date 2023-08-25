package n41

import (
	"container/heap"
	"fmt"
	"n41/n41msg"
	"n41/n41types"
	"sync"
	"time"
)

const (
	REQUEST_TIMEOUT             = int64(300 * time.Millisecond)
	REQUEST_SENDING_RETRY_COUNT = uint8(2)
	RESPONSE_RETENTION          = int64(1000 * time.Millisecond)
	EXPIRING_CHECK_INTERVAL     = 10 * time.Millisecond
)

type SendingInfo struct {
	msg        *n41msg.Message
	remote     *n41types.Sbi // must be non-nil
	expire     int64         // expiring time
	onexpiring func()
	index      int // index that will be use to build a priority queue
}

func newSendingInfo(msg *n41msg.Message, remote *n41types.Sbi, dur int64, onexpiring func()) (info SendingInfo) {
	info.msg = msg
	info.remote = remote
	info.expire = time.Now().UnixNano() + dur
	info.onexpiring = onexpiring
	return
}

func (info *SendingInfo) DeadTime() int64 {
	return info.expire
}

func (info *SendingInfo) Id() ItemId {
	return ItemId{
		remote: info.remote.String(),
	}
}

func (info *SendingInfo) SetIndex(i int) {
	info.index = i
}

func (info *SendingInfo) Index() int {
	return info.index
}
func (info *SendingInfo) OnExpiring() func() {
	return info.onexpiring
}

type ReqSendingInfo struct {
	SendingInfo
	err   error
	rsp   *n41msg.Message //response (for casting)
	retry uint8           //counter of retrying
	done  chan bool       //close when receiving a response (or timer expiring)
}

func newReqSendingInfo(msg *n41msg.Message, remote *n41types.Sbi, scheduler func(*ReqSendingInfo)) (info *ReqSendingInfo) {
	info = &ReqSendingInfo{
		done:  make(chan bool),
		retry: 0,
	}
	info.SendingInfo = newSendingInfo(msg, remote, REQUEST_TIMEOUT, func() {
		if info.retry < REQUEST_SENDING_RETRY_COUNT {
			//resend
			info.expire += REQUEST_TIMEOUT
			scheduler(info)
			// Note: there is an neglicible possibility that a response arrives just before the request is re-schedule for sending. In such a case the response will be ignored (as its corresponding request not found. Still, the implementation is safe!
		} else {
			info.err = fmt.Errorf("Request timeout error")
			close(info.done)
		}
	})
	return
}

type RspSendingInfo struct {
	SendingInfo
}

func newRspSendingInfo(msg *n41msg.Message, remote *n41types.Sbi) *RspSendingInfo {
	return &RspSendingInfo{
		newSendingInfo(msg, remote, RESPONSE_RETENTION, nil),
	}
}

// for indexing
type ItemId struct {
	remote string
}

type ExpiringItem interface {
	DeadTime() int64
	Id() ItemId
	SetIndex(int)
	Index() int
	OnExpiring() func()
}

// priority queue
type PriorityQueue []ExpiringItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].DeadTime() < pq[j].DeadTime()
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].SetIndex(i)
	pq[j].SetIndex(j)
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(ExpiringItem)
	item.SetIndex(n)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

type ExpiringList struct {
	id2item map[ItemId]ExpiringItem
	pq      PriorityQueue
	mux     sync.Mutex
}

func newExpiringList() (l ExpiringList) {
	l = ExpiringList{
		id2item: make(map[ItemId]ExpiringItem),
	}
	//heap.Init(&l.pq)
	return l
}

func (l *ExpiringList) find(remote string, seq uint32) (item ExpiringItem) {
	l.mux.Lock()
	defer l.mux.Unlock()
	id := ItemId{
		remote: remote,
		// seq:    seq,
	}
	item, _ = l.id2item[id]
	return
}

func (l *ExpiringList) pop(remote string) (item ExpiringItem) {
	l.mux.Lock()
	defer l.mux.Unlock()
	id := ItemId{
		remote: remote,
	}
	var ok bool
	if item, ok = l.id2item[id]; ok {
		//logrus.Info("pop a found item from list")
		delete(l.id2item, id)
		heap.Remove(&l.pq, item.Index())
	}
	return
}

func (l *ExpiringList) add(item ExpiringItem) {
	l.flush()

	l.mux.Lock()
	defer l.mux.Unlock()
	// logrus.Info("add item to list")
	l.id2item[item.Id()] = item
	heap.Push(&l.pq, item)
}

// remove expring items
func (l *ExpiringList) flush() {
	l.mux.Lock()
	defer l.mux.Unlock()
	t := time.Now().UnixNano()
	//count the number of expiring items
	cnt := 0
	for i := len(l.pq) - 1; i >= 0; i-- {
		if l.pq[i].DeadTime() > t {
			break
		}
		cnt++
	}

	//remove expiring items and fire their callbacks
	if cnt > 0 {
		for i := 0; i < cnt; i++ {
			//pop from the priority queue
			item := heap.Pop(&l.pq).(ExpiringItem)
			//logrus.Info("remove item from list due to timing out")
			//delete from the map
			delete(l.id2item, item.Id())
			if fn := item.OnExpiring(); fn != nil {
				fn()
			}

		}
	}
	/*
		//simple implementation without a priority queue (will be removed)
			for id, item := range l.list {
				if t >= item.DeadTime() {
					delete(l.id2item, id)
					if fn := item.OnExpiring(); fn != nil {
						fn()
					}
				}
			}
	*/
}
