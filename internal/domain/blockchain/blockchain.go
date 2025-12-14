package blockchain

import (
	"github.com/Mohsen20031203/learn-gochain-core/internal/domain/block"
)

type Blockchain struct {
	Difficulty int
}

func New() *Blockchain {
	return &Blockchain{
		Difficulty: 3,
	}
}

func (bc *Blockchain) ValidateNewBlock(prev, new block.Block) bool {
	if new.PrevHash != prev.Hash {
		return false
	}
	if new.CalculateHash() != new.Hash {
		return false
	}
	return new.HasValidPoW(bc.Difficulty)
}
