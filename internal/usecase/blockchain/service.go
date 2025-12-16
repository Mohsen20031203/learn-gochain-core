package blockchain

import (
	"errors"
	"time"

	"github.com/Mohsen20031203/learn-gochain-core/config"
	"github.com/Mohsen20031203/learn-gochain-core/internal/domain/block"
	"github.com/Mohsen20031203/learn-gochain-core/internal/domain/node"
	"github.com/Mohsen20031203/learn-gochain-core/internal/infrastructure/storage/lvldb"
	"github.com/syndtr/goleveldb/leveldb"
)

// 1. validata the tx
// 2. put the mempool
// 3. mine a new block
// 4. put the txs mompool to the new block
// 5. validate the block and txs
// 6. save new block to the repo

type NodeService struct {
	node   *node.Node
	repo   Repository
	config config.Config
}

func NewService(config config.Config) *NodeService {

	repo := lvldb.New(config.FileStoragePath)
	repo.Open()

	node := node.NewNode(config.NodeID, config.Difficulty)

	return &NodeService{
		node:   node,
		repo:   repo,
		config: config,
	}
}

func (s *NodeService) AddBlock(data string) (*block.Block, error) {
	last, err := s.repo.Get(LastBlockKey)
	if err != leveldb.ErrNotFound && err != nil {
		return nil, err
	}

	if last != nil {
		if !last.HasValidPoW(s.node.Chain.Difficulty) {
			return nil, errors.New("last block has invalid proof of work")
		}
	}
	var ind int64 = 0

	if last != nil {
		ind = last.Index() + 1
		newBlock.Index() = last.Index + 1
		newBlock.PrevHash = last.Hash
	}

	newBlock := block.Block{
		Timestamp: time.Now(),
		Data:      data,
		Index:     ind,
		PrevHash:  "0",
	}

	newBlock.Mine(s.node.Chain.Difficulty)

	if last != nil && !s.node.Chain.ValidateNewBlock(last, &newBlock) {
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

func (s *NodeService) GetChain() ([]block.Block, error) {
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
