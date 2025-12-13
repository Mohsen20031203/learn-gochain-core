package blockchain

import (
	"time"

	"github.com/Mohsen20031203/learn-gochain-core/internal/domain/block"
	"github.com/Mohsen20031203/learn-gochain-core/internal/domain/blockchain"
)

type Service struct {
	chain *blockchain.Blockchain
}

func NewService(chain *blockchain.Blockchain) *Service {
	return &Service{chain: chain}
}

func (s *Service) CreateBlock(b block.Block) (block.Block, error) {
	b.Timestamp = time.Now()

	err := s.chain.AddBlock(&b)
	if err != nil {
		return block.Block{}, err
	}

	return b, nil
}

func (s *Service) GetChain() []block.Block {
	return s.chain.Blocks
}
