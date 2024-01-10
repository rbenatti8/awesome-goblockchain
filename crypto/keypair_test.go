package crypto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKeypairSignVerifySuccess(t *testing.T) {
	privateKey := GeneratePrivateKey()

	msg := []byte("Hello, world!")
	signature, err := privateKey.Sign(msg)

	assert.Nil(t, err)
	assert.True(t, signature.Verify(msg, privateKey.PublicKey()))
}

func TestKeypairSignVerifyFail(t *testing.T) {
	privateKey := GeneratePrivateKey()

	msg := []byte("Hello, world!")
	signature, err := privateKey.Sign(msg)

	otherPrivateKey := GeneratePrivateKey()

	assert.Nil(t, err)
	assert.False(t, signature.Verify(msg, otherPrivateKey.PublicKey()))
	assert.False(t, signature.Verify([]byte("Hello, world"), privateKey.PublicKey()))
}
