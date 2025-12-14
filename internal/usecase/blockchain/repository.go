package blockchain

import (
	"github.com/Mohsen20031203/learn-gochain-core/internal/domain/block"
)

const LastBlockKey = "LastBlock"

type Repository interface {
	Open() error
	Save(key string, data *block.Block) error
	Load() (*block.Block, error)
	Get(key string) (*block.Block, error)
	Close() error
}
