package block

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Index     int64     `json:"index"`
	Timestamp time.Time `json:"timestamp"`
	Data      string    `json:"data"`
	Hash      string    `json:"hash"`
	PrevHash  string    `json:"prev_hash"`
	Nonce     int64     `json:"nonce"`
}

func (b *Block) CalculateHash() string {
	record :=
		strconv.FormatInt(b.Index, 10) +
			b.Timestamp.String() +
			b.Data +
			b.PrevHash +
			strconv.FormatInt(b.Nonce, 10)

	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}

func (b *Block) Mine(difficulty int) {
	prefix := strings.Repeat("0", difficulty)

	for {
		hash := b.CalculateHash()
		if strings.HasPrefix(hash, prefix) {
			b.Hash = hash
			break
		}
		b.Nonce++
	}
}

func (b *Block) IsValid(prev Block) bool {
	if b.PrevHash != prev.Hash {
		return false
	}
	if b.CalculateHash() != b.Hash {
		return false
	}
	return true
}

func (b *Block) HasValidPoW(difficulty int) bool {
	prefix := strings.Repeat("0", difficulty)
	return strings.HasPrefix(b.Hash, prefix)
}
