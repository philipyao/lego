package network

import (
	"net"
	"sync"
	"time"
)

type NewSession func(SessionConn) TCPSession

type TCPServer struct {
	Addr    string
	ln      net.Listener

	NewSession  NewSession
	wg      sync.WaitGroup

	ReadTimeout     int
	WriteTimeout    int
}

func NewTCPServer(addr string, f NewSession) *TCPServer {
	server := new(TCPServer)
	server.Addr = addr
	server.NewSession = f

	return server
}

func (ts *TCPServer) Serve() {
	err := ts.init()
	if err != nil {
		panic(err)
	}

	go ts.run()
}

func (ts *TCPServer) Close() {
	if ts.ln != nil  {
		ts.ln.Close()
	}
}

//==============================================

func (ts *TCPServer) init() error {
	l, err := net.Listen("tcp", ts.Addr)
	if err != nil {
		return err
	}
	ts.ln = l
	return nil
}

func (ts *TCPServer) run() error {
	for {
		conn, err := ts.ln.Accept()
		if err != nil {
			//log
			break
		}

        //TODO SetReadBuffer/SetWriteBuffer

		sc := SessionConn{
			conn: conn,
			ts: ts,

			sendChan: make(chan []byte, 100),
			recvChan: make(chan interface{}, 100),

			readTmo: time.Duration(5),
			writeTmo: time.Duration(5),
		}
		sess := ts.NewSession(sc)
		err = sess.OnOpen()
		if err != nil {

		}

		//收包
		go func() {
			for {
				msg, err := sess.ReadMsg()
				if err != nil {

				}
				err = sess.HandleMsg(msg)
				if err != nil {

				}
			}
		}()

		//发包
		go sess.HandleSendMsg()
	}
	return nil
}
