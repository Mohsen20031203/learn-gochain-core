package block

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type Block struct {
	index     int64
	timestamp time.Time
	data      string
	hash      string
	prevHash  string
	nonce     int64
}

func NewBlock(index int64, data string, prevHash string) *Block {
	block := &Block{
		index:     index,
		timestamp: time.Now(),
		data:      data,
		prevHash:  prevHash,
		nonce:     0,
	}
	return block
}

func (b *Block) Index() int64 {
	return b.index
}

func (b *Block) Timestamp() time.Time {
	return b.timestamp
}

func (b *Block) Data() string {
	return b.data
}

func (b *Block) Hash() string {
	return b.hash
}

func (b *Block) PrevHash() string {
	return b.prevHash
}

func (b *Block) Nonce() int64 {
	return b.nonce
}

func (b *Block) SetHash(hash string) {
	b.hash = hash
}

func (b *Block) SetPrevHash(prevHash string) {
	b.prevHash = prevHash
}

func (b *Block) SetNonce(nonce int64) {
	b.nonce = nonce
}

func (b *Block) CalculateHash() string {
	record :=
		strconv.FormatInt(b.index, 10) +
			b.timestamp.String() +
			b.data +
			b.prevHash +
			strconv.FormatInt(b.nonce, 10)

	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}

func (b *Block) IsValid(prev Block) bool {
	if b.prevHash != prev.hash {
		return false
	}
	if b.CalculateHash() != b.hash {
		return false
	}
	return true
}
