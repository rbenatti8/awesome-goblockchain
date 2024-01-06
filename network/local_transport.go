package network

import (
	"fmt"
	"sync"
)

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

func (t *LocalTransport) Consume() <-chan RPC {
	return t.consumeChan
}

func (t *LocalTransport) Connect(transport Transport) error {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.peers[transport.Addr()] = transport.(*LocalTransport)
	return nil
}

func (t *LocalTransport) Addr() NetAddr {
	return t.addr
}

func (t *LocalTransport) SendMessage(addr NetAddr, payload []byte) error {
	t.lock.RLock()
	defer t.lock.RUnlock()
	peer, ok := t.peers[addr]
	if !ok {
		return fmt.Errorf("peer %s not found", addr)
	}

	peer.consumeChan <- RPC{
		From:    t.addr,
		Payload: payload,
	}

	return nil
}
