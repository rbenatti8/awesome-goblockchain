package core

import (
	"awesome-goblockchain/types"
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestHeader_EncodeBinary(t *testing.T) {
	h := Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		Timestamp: time.Now().UnixNano(),
		Height:    1,
		Nonce:     9898767,
	}

	b := new(bytes.Buffer)

	assert.NoError(t, h.EncodeBinary(b))
}

func TestHeader_DecodeBinary(t *testing.T) {
	h := Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		Timestamp: time.Now().UnixNano(),
		Height:    1,
		Nonce:     9898767,
	}

	b := new(bytes.Buffer)

	assert.NoError(t, h.EncodeBinary(b))

	var h2 Header

	assert.NoError(t, h2.DecodeBinary(b))

	assert.Equal(t, h, h2)
}

func TestBlock_EncodeBinary(t *testing.T) {
	h := Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		Timestamp: time.Now().UnixNano(),
		Height:    1,
		Nonce:     9898767,
	}

	block := Block{
		Header:       h,
		Transactions: nil,
	}

	b := new(bytes.Buffer)

	assert.NoError(t, block.EncodeBinary(b))

}

func TestBlock_DecodeBinary(t *testing.T) {
	h := Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		Timestamp: time.Now().UnixNano(),
		Height:    1,
		Nonce:     9898767,
	}

	block := Block{
		Header:       h,
		Transactions: nil,
	}

	b := new(bytes.Buffer)

	assert.NoError(t, block.EncodeBinary(b))

	var block2 Block

	assert.NoError(t, block2.DecodeBinary(b))

	assert.Equal(t, block, block2)
}

func TestBlock_Hash(t *testing.T) {
	h := Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		Timestamp: time.Now().UnixNano(),
		Height:    1,
		Nonce:     9898767,
	}

	block := Block{
		Header:       h,
		Transactions: nil,
	}

	hash := block.Hash()
	assert.False(t, hash.IsZero())

	assert.Equal(t, hash, block.Hash())
}
