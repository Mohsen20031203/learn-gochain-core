package blockchain

import (
	"errors"
	"time"

	"github.com/Mohsen20031203/learn-gochain-core/config"
	"github.com/Mohsen20031203/learn-gochain-core/internal/domain/block"
	"github.com/Mohsen20031203/learn-gochain-core/internal/domain/blockchain"
	"github.com/Mohsen20031203/learn-gochain-core/internal/infrastructure/storage/lvldb"
	"github.com/syndtr/goleveldb/leveldb"
)

type Service struct {
	chain  *blockchain.Blockchain
	repo   Repository
	config config.Config
}

func NewService(config config.Config) *Service {

	repo := lvldb.New(config.FileStoragePath)
	repo.Open()

	chain := blockchain.Blockchain{
		Difficulty: config.Difficulty,
	}

	return &Service{chain: &chain, repo: repo, config: config}
}

func (s *Service) AddBlock(data string) (*block.Block, error) {
	last, err := s.repo.Get(LastBlockKey)
	if err != leveldb.ErrNotFound && err != nil {
		return nil, err
	}

	if last != nil {
		if !last.HasValidPoW(s.chain.Difficulty) {
			return nil, errors.New("last block has invalid proof of work")
		}
	}

	newBlock := block.Block{
		Timestamp: time.Now(),
		Data:      data,
		Index:     0,
		PrevHash:  "0",
	}

	if last != nil {
		newBlock.Index = last.Index + 1
		newBlock.PrevHash = last.Hash
	}

	newBlock.Mine(s.chain.Difficulty)

	if last != nil && !s.chain.ValidateNewBlock(*last, newBlock) {
		return nil, errors.New("invalid block")
	}

	if last != nil {
		if err := s.repo.Save(last.Hash, last); err != nil {
			return nil, err
		}
	}

	if err := s.repo.Save(LastBlockKey, &newBlock); err != nil {
		return nil, err
	}

	return &newBlock, nil
}

func (s *Service) GetChain() ([]block.Block, error) {
	var chain []block.Block

	current, err := s.repo.Get(LastBlockKey)
	if err != nil {
		return nil, err
	}
	for current != nil {
		chain = append([]block.Block{*current}, chain...) // prepend
		if current.PrevHash == "0" {
			break
		}
		current, err = s.repo.Get(current.PrevHash)
		if err != nil {
			return nil, err
		}
	}

	return chain, nil
}
