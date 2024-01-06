package network

import (
	"fmt"
	"time"
)

type ServerOps struct {
	Transports []Transport
}

type Server struct {
	ops      ServerOps
	rpcChan  chan RPC
	quitChan chan struct{}
}

func NewServer(ops ServerOps) *Server {
	return &Server{
		ops:      ops,
		rpcChan:  make(chan RPC),
		quitChan: make(chan struct{}),
	}
}

func (s *Server) Start() {
	s.initTransports()
	ticker := time.NewTicker(5 * time.Second)

	for {
		select {
		case rpc := <-s.rpcChan:
			fmt.Printf("RPC: %s\n", rpc.Payload)
		case <-s.quitChan:
			fmt.Println("Server stopped")
			return
		case <-ticker.C:
			fmt.Println("Server is running")
		}
	}
}

func (s *Server) initTransports() {
	for _, transport := range s.ops.Transports {
		go func(tr Transport) {
			for rpc := range tr.Consume() {
				s.rpcChan <- rpc
			}
		}(transport)
	}
}
