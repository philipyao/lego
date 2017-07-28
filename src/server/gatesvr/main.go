package main

import (
    "time"
    "components/network"
)

var server *network.TCPServer

func main() {
    server = network.NewTCPServer("localhost:4565", func(c network.SessionConn) network.TCPSession{
        return &GateSession{
	        SessionConn: c,
	        Foo: 0,
	        Bar: "hello,foobar",
        }
    }) 
    server.Serve()

    time.Sleep(time.Duration(100) * time.Second)

    server.Close()
}
