package block

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

type Block struct {
	Index     int64
	Timestamp time.Time
	Data      string `json:"data"`
	//Transactions []Transaction
	PrevHash string
	Hash     string
	Nonce    int64
}

func (b Block) CalculateHash() string {
	record := fmt.Sprintf("%d%s%s%d", b.Index, b.Timestamp.String(), b.PrevHash, b.Nonce)
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
