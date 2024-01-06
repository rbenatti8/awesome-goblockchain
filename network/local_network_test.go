package network

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Path: local_network_test.go

func TestNewLocalTransport(t *testing.T) {
	t.Parallel()

	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	_ = tra.Connect(trb)
	_ = trb.Connect(tra)

	assert.Equal(t, tra.peers[trb.addr], trb)
	assert.Equal(t, trb.peers[tra.addr], tra)
}

func TestLocalTransport_SendMessage(t *testing.T) {
	t.Parallel()

	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	_ = tra.Connect(trb)
	_ = trb.Connect(tra)

	assert.Nil(t, tra.SendMessage(trb.addr, []byte("Hello World!")))

	rpc := <-trb.Consume()

	assert.Equal(t, rpc.From, tra.addr)
	assert.Equal(t, rpc.Payload, []byte("Hello World!"))
}
