package crypto

import (
	"awesome-goblockchain/types"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"math/big"
)

type PrivateKey struct {
	key *ecdsa.PrivateKey
}

func (k *PrivateKey) Sign(msg []byte) (*Signature, error) {
	r, s, err := ecdsa.Sign(rand.Reader, k.key, msg)
	if err != nil {
		return nil, err
	}

	return &Signature{
		s: s,
		r: r,
	}, nil
}

func (k *PrivateKey) PublicKey() *PublicKey {
	return &PublicKey{
		key: &k.key.PublicKey,
	}
}

func GeneratePrivateKey() *PrivateKey {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	return &PrivateKey{
		key: key,
	}
}

type PublicKey struct {
	key *ecdsa.PublicKey
}

func (k *PublicKey) toBytes() []byte {
	return elliptic.MarshalCompressed(elliptic.P256(), k.key.X, k.key.Y)
}
func (k *PublicKey) Address() types.Address {
	h := sha256.Sum256(k.toBytes())
	return types.Address(h[len(h)-20:])
}

type Signature struct {
	s, r *big.Int
}

func (s Signature) Verify(msg []byte, pub *PublicKey) bool {
	return ecdsa.Verify(pub.key, msg, s.r, s.s)
}
