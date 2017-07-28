package main

import (
	"components/network"
)

type GateSession struct {
	network.SessionConn

	Foo     int
	Bar     string
}

func (gs *GateSession) OnOpen() error {
	return nil
}

func (gs *GateSession) ReadMsg() (interface{}, error) {
	return nil, nil
}

func (gs *GateSession) WriteMsg(msg interface{}) error {
	return nil
}

func (gs *GateSession) HandleMsg(msg interface{}) error {
	return nil
}

func (gs *GateSession) OnClose() error {
	return nil
}
