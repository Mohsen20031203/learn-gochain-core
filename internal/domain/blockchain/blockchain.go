package blockchain

import (
	"errors"

	"github.com/Mohsen20031203/learn-gochain-core/internal/domain/block"
)

type Blockchain struct {
	Blocks     []block.Block
	Difficulty int
}

func New() *Blockchain {
	return &Blockchain{
		Blocks:     []block.Block{},
		Difficulty: 6,
	}
}

func (bc *Blockchain) AddBlock(b *block.Block) error {
	if len(bc.Blocks) == 0 {
		b.Index = 0
		b.PrevHash = "0"
	} else {
		prev := bc.Blocks[len(bc.Blocks)-1]
		b.Index = int64(len(bc.Blocks))
		b.PrevHash = prev.Hash
	}

	b.Mine(bc.Difficulty)

	if len(bc.Blocks) > 0 {
		prev := bc.Blocks[len(bc.Blocks)-1]
		if !b.IsValid(prev) {
			return errors.New("invalid block")
		}
	}

	bc.Blocks = append(bc.Blocks, *b)
	return nil
}
