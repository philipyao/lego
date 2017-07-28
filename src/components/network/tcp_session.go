package network

import (
	"net"
	"time"
)

type TCPSession interface {
	OnOpen() error
	ReadMsg() (interface{}, error)
	WriteMsg(interface{}) error
	HandleMsg(interface{}) error
	OnClose() error

	RemoteAddr() string
	HandleSendMsg()
}

// session连接
type SessionConn struct {
	conn net.Conn
	ts   *TCPServer

	sendChan chan []byte
	recvChan chan interface{}

	readTmo time.Duration
	writeTmo time.Duration
}

func (sc *SessionConn) RemoteAddr() string {
	return sc.conn.RemoteAddr().String()
}

func (sc *SessionConn) HandleSendMsg() {
	for b := range sc.sendChan {
		if b == nil {
			break
		}
		sc.conn.SetWriteDeadline(time.Now().Add(sc.writeTmo))
		l, err := sc.conn.Write(b)
		if err != nil {
			//log
			break
		}
		if l != len(b) {
			//log
			break
		}
	}
}