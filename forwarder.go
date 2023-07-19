package n41

import (
	"fmt"
	"n41/n41msg"
	"n41/n41types"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	N41_BUF_SIZE int = 2048
)

type RecvInfo struct {
	msg *n41msg.Message
	// remote *net.UDPAddr
	remote *n41types.Sbi
}

type Forwarder struct {
	// conn *net.UDPConn
	// addr net.UDPAddr
	conn *http.Server
	addr n41types.Sbi
	when time.Time //started time
	wg   sync.WaitGroup
}

// func newForwarder(addr net.UDPAddr) *Forwarder {
func newForwarder(addr n41types.Sbi) *Forwarder {
	ret := &Forwarder{
		addr: addr,
	}
	return ret
}

func (fwd *Forwarder) start(recv chan<- RecvInfo) (err error) {

	// if fwd.conn, err = net.ListenUDP("udp", &fwd.addr); err != nil {
	// 	logrus.Errorf("Failed to listen: %s", err.Error())
	// 	return
	// }

	// logrus.Infof("Listen on N4 interface %s", fwd.conn.LocalAddr().String())

	gin.SetMode(gin.ReleaseMode)
	route := gin.Default()

	fwd.conn = &http.Server{
		Addr:    fwd.addr.GetAddr(),
		Handler: route,
	}
	go func() {
		err := fwd.conn.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			fmt.Println("Server error:", err)
		}
	}()

	// go fwd.loop(recv)
	fwd.when = time.Now()
	return
}

func (fwd *Forwarder) loop(recv chan<- RecvInfo) {
	fwd.wg.Add(1)
	defer fwd.wg.Done()
	buf := make([]byte, N41_BUF_SIZE)
	var msg *n41msg.Message
	for {
		if n, addr, err := fwd.conn.ReadFromUDP(buf); err == nil {
			msg = new(n41msg.Message)
			if err = msg.Unmarshal(buf[:n]); err == nil {
				if recv != nil {
					recv <- RecvInfo{
						msg:    msg,
						remote: addr,
					}
				}
			}
		} else {
			logrus.Errorf(err.Error())
			break
		}
	}
	if recv != nil {
		close(recv)
	}
}
func (fwd *Forwarder) stop() {
	fwd.conn.Close()
	fwd.wg.Wait()
}

//time when the forwarder started running
func (fwd *Forwarder) When() time.Time {
	return fwd.when
}

// block util message is written to the transport or an error occurs
// func (fwd *Forwarder) WriteTo(msg *n41msg.Message, addr *net.UDPAddr) (err error) {
func (fwd *Forwarder) WriteTo(msg *n41msg.Message, addr *n41types.Sbi) (err error) {
	var buf []byte
	if buf, err = msg.Marshal(); err == nil {
		_, err = fwd.conn.WriteToUDP(buf, addr)
	}
	return
}
