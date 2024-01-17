package network

import "sync"

type LocalTransport struct {
	addr        NetAddr
	consumeChan chan RPC
	lock        sync.RWMutex
	peers       map[NetAddr]*LocalTransport
}

func NewLocalTransport(addr NetAddr) *LocalTransport {
	return &LocalTransport{
		addr:        addr,
		consumeChan: make(chan RPC, 1024),
		peers:       make(map[NetAddr]*LocalTransport),
	}
}
