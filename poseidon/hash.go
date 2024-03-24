package poseidon

import (
	"hash"

	"github.com/iden3/go-iden3-crypto/poseidon"
)

const FRAME_SIZE = 16

type PoseidonHash struct {
	inner hash.Hash
}

func New() *PoseidonHash {
	p, _ := poseidon.New(FRAME_SIZE)
	return &PoseidonHash{p}
}

func (ph *PoseidonHash) Hash(data ...[]byte) []byte {
	var hash []byte
	if len(data) == 1 {
		hash = poseidon.Sum(data[0])
	} else {
		concatDataLen := 0
		for _, d := range data {
			concatDataLen += len(d)
		}
		concatData := make([]byte, concatDataLen)
		curOffset := 0
		for _, d := range data {
			copy(concatData[curOffset:], d)
			curOffset += len(d)
		}
		hash = poseidon.Sum(concatData)
	}

	return hash
}

func (ph *PoseidonHash) HashLength() int {
	return ph.inner.Size()
}

func (ph *PoseidonHash) HashName() string {
	return "poseidon"
}
