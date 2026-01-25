package blockchain

import (
	"context"
	"strings"

	"github.com/Mohsen20031203/learn-gochain-core/internal/domain/block"
)

type Blockchain struct {
	difficulty int
	lastHash   string
	// Count of blocks in the chain
	height int
}

func New(difficulty int) *Blockchain {
	chain := Blockchain{}

	chain.setDifficulty(difficulty)
	return &chain
}

func (bc *Blockchain) GetDifficulty() int {
	return bc.difficulty
}

func (bc *Blockchain) setDifficulty(difficulty int) {
	bc.difficulty = difficulty
}

func (bc *Blockchain) IsValidNewBlock(block *block.Block) bool {
	return block.PrevHash == bc.lastHash
}

func (bc *Blockchain) UpdateWithNewBlock(block *block.Block) {
	bc.setLastHash(block.Hash)
	bc.height++
}

func (bc *Blockchain) setLastHash(hash string) {
	bc.lastHash = hash
}

func (bc *Blockchain) LastBlockHash() string {
	return bc.lastHash
}

func (bc *Blockchain) CountBlocks() int {
	return bc.height
}

func (bc *Blockchain) Mine(ctx context.Context, b *block.Block) bool {
	prefix := strings.Repeat("0", bc.difficulty)

	for {
		select {
		case <-ctx.Done():
			return false
		default:
		}

		hash := b.CalculateHash()
		if strings.HasPrefix(hash, prefix) {
			b.Hash = hash
			return true
		}
		b.Nonce++
	}
}

func (bc *Blockchain) IsValidPoW(b *block.Block) bool {
	prefix := strings.Repeat("0", bc.difficulty)
	return strings.HasPrefix(b.Hash, prefix)
}
