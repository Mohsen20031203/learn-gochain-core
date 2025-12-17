package blockchain

import (
	"context"
	"fmt"

	"github.com/Mohsen20031203/learn-gochain-core/config"
	"github.com/Mohsen20031203/learn-gochain-core/internal/domain/block"
	"github.com/Mohsen20031203/learn-gochain-core/internal/domain/node"
	"github.com/Mohsen20031203/learn-gochain-core/internal/domain/transaction"
	"github.com/Mohsen20031203/learn-gochain-core/internal/infrastructure/storage/lvldb"
)

// 1. validata the tx
// 2. put the mempool
// 3. mine a new block
// 4. put the txs mompool to the new block
// 5. validate the block and txs
// 6. save new block to the repo

type NodeService struct {
	node        *node.Node
	repo        Repository
	config      config.Config
	mineTrigger chan struct{}
}

func NewService(config config.Config) *NodeService {

	repo := lvldb.New(config.FileStoragePath)
	repo.Open()

	node := node.NewNode(config.NodeID, config.Difficulty)

	return &NodeService{
		node:        node,
		repo:        repo,
		config:      config,
		mineTrigger: make(chan struct{}),
	}
}

func (s *NodeService) SubmitTransactions(tx []transaction.Transaction) error {

	for _, t := range tx {
		s.node.AddTransactionMempool(t)
	}
	if s.node.SizeMempool() >= s.config.BatchSize {
		select {
		case s.mineTrigger <- struct{}{}:
		default:
		}
	}
	return nil
}

func (s *NodeService) GetLastBlock() (*block.Block, error) {
	last, err := s.repo.Get(LastBlockKey)
	if err != nil {
		return nil, err
	}
	if last.Hash == "" {
		return nil, nil
	}
	return last, nil
}

func (s *NodeService) StartMiner(ctx context.Context) {
	go func() {
		for {
			select {
			case <-s.mineTrigger:
				for {
					if s.node.SizeMempool() < s.config.BatchSize {
						break
					}
					s.mineOnce()
				}
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (s *NodeService) mineOnce() {
	tx2 := s.node.GetMempoolTransaction(s.config.BatchSize)
	if len(tx2) == 0 {
		return
	}

	tx := make([]transaction.Transaction, len(tx2))
	copy(tx, tx2)

	lastBlock := s.node.GetChainLastBlockHash()

	var blc *block.Block
	if lastBlock == "" {
		blc = block.NewBlock(0, tx, "0")
	} else {
		blc = block.NewBlock(s.node.CountBlocksinChain(), tx, lastBlock)
	}

	s.node.MineBlock(blc)

	if lastBlock != "" && !s.node.IsValidNewBlockChain(*blc) {
		fmt.Println("Invalid mined block")
		return
	}

	s.saveBlock(blc)

	if s.node.SizeMempool() == len(tx) {
		s.node.ClearMempool()
		return
	}

	for _, t := range tx {
		s.node.RemoveTransactionMempool(t)
	}
}

func (s *NodeService) saveBlock(b *block.Block) error {
	lastBlock := s.node.GetChainLastBlockHash()
	if lastBlock != "" {
		bl, err := s.repo.Get(lastBlock)
		if err != nil {
			return err
		}
		if err := s.repo.Save(lastBlock, bl); err != nil {
			return err
		}
	}

	if err := s.repo.Save(LastBlockKey, b); err != nil {
		return err
	}
	if err := s.repo.Save(b.Hash, b); err != nil {
		return err
	}
	s.node.UpdateChain(*b)
	return nil
}

func (s *NodeService) GetMempoolTransactions() []transaction.Transaction {
	return s.node.GetMempoolTransactions()
}

func (s *NodeService) GetChain() ([]block.Block, error) {
	var chain []block.Block

	current, err := s.repo.Get(LastBlockKey)
	if err != nil {
		return nil, err
	}
	for current.Hash != "" {
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
