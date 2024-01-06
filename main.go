package main

import (
	"awesome-goblockchain/network"
	"time"
)

func main() {
	trLocal := network.NewLocalTransport("LOCAL")
	trRemote := network.NewLocalTransport("REMOTE")

	trLocal.Connect(trRemote)
	trRemote.Connect(trLocal)

	go func() {
		for {
			trRemote.SendMessage(trLocal.Addr(), []byte("Hello World!"))
			time.Sleep(1 * time.Second)
		}

	}()
	serv := network.NewServer(network.ServerOps{
		Transports: []network.Transport{
			trLocal,
			trRemote,
		},
	})

	serv.Start()
}
