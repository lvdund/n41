package n41

import (
	"bytes"
	"encoding/json"
	"fmt"
	"n41/n41msg"
	"n41/n41types"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	N41_BUF_SIZE int = 2048
)

type RecvInfo struct {
	msg    *n41msg.Message
	remote *n41types.SbiAdrr // address of requester
}

type Forwarder struct {
	// conn *net.UDPConn
	// addr net.UDPAddr
	conn *http.Server
	addr n41types.SbiAdrr
	when time.Time //started time
	// wg   sync.WaitGroup
}

func newForwarder(addr n41types.SbiAdrr) *Forwarder {
	ret := &Forwarder{
		addr: addr,
	}
	return ret
}

func (fwd *Forwarder) start(recv chan<- RecvInfo) (err error) {

	gin.SetMode(gin.ReleaseMode)
	route := gin.Default()

	route.GET("/", func(ctx *gin.Context) {
		var msg *n41msg.Message
		rawData, err := ctx.GetRawData()
		if err != nil {
			return
		}
		if err = msg.N41Unmarshal(rawData); err == nil {
			if recv != nil {
				recv <- RecvInfo{
					msg:    msg,
					remote: &msg.Header.Adrr,
				}
			}
		} else {
			logrus.Errorf(err.Error())
		}
		// if recv != nil {
		// 	close(recv)
		// }
	})

	fwd.conn = &http.Server{
		Addr:    fwd.addr.GetAddr(),
		Handler: route,
	}
	// go func() {
	// 	err := fwd.conn.ListenAndServe()
	// 	if err != nil && err != http.ErrServerClosed {
	// 		fmt.Println("Server error:", err)
	// 	}
	// }()
	route.Run(fwd.addr.GetAddr())

	logrus.Infof("Listen on N4 interface %s", fwd.conn.Addr)

	fwd.when = time.Now()
	return
}

func (fwd *Forwarder) stop() {
	fwd.conn.Close()
	// fwd.wg.Wait()
}

//time when the forwarder started running
func (fwd *Forwarder) When() time.Time {
	return fwd.when
}

// block util message is written to the transport or an error occurs
func (fwd *Forwarder) WriteTo(msg *n41msg.Message, addr *n41types.SbiAdrr) (err error) {
	sendMsg, err := json.Marshal(msg)
	url := url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%d", addr.IP, addr.Port),
	}
	req, err := http.NewRequest(http.MethodGet, url.String(), bytes.NewBuffer(sendMsg))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	_, err = client.Do(req)

	return
}
