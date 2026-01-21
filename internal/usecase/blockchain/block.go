package blockchain

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Mohsen20031203/learn-gochain-core/internal/domain/block"
	"github.com/Mohsen20031203/learn-gochain-core/internal/domain/transaction"
	"github.com/Mohsen20031203/learn-gochain-core/internal/infrastructure/network"
)

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

func (s *NodeService) validataBlock(blc block.Block) bool {

	lastBlockHash := s.node.GetChainLastBlockHash()
	if lastBlockHash != "" {
		if !s.node.IsValidNewBlockChain(blc) {
			fmt.Println("Invalid block: previous hash does not match")
			return false
		}
	} else {
		if blc.Index != 0 {
			fmt.Println("Invalid block: genesis block index must be 0")
			return false
		}
		if !s.node.IsValidPoW(&blc) {
			fmt.Println("Invalid block: proof of work is not valid")
			return false
		}
		return true
	}

	if blc.Index < s.node.CountBlocksinChain() {
		fmt.Println("Invalid block: index is not greater than last block index")
		return false
	}

	if !s.node.IsValidPoW(&blc) {
		fmt.Println("Invalid block: proof of work is not valid")
		return false
	}

	b, err := s.repo.Get(blc.Hash)
	if err == nil && b.Hash != "" {
		fmt.Println("Invalid block: block already exists")
		return false
	}

	return true
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
	blc.Timestamp = time.Now()

	if lastBlock != "" && !s.node.IsValidNewBlockChain(*blc) {
		fmt.Println("Invalid mined block")
		return
	}

	s.saveBlock(blc)
	s.broadcastBlock(blc)

	if s.node.SizeMempool() == len(tx) {
		s.node.ClearMempool()
		return
	}

	for _, t := range tx {
		s.node.RemoveTransactionMempool(t)
	}
}

func (s *NodeService) broadcastBlock(blc *block.Block) {
	if s.broadcaster == nil {
		return
	}

	data, err := json.Marshal(blc)
	if err != nil {
		return
	}

	msg := network.Message{
		Type: "block",
		Data: data,
	}

	s.broadcaster.Broadcast(msg)
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

func (s *NodeService) GetBlockByHash(block string) (*block.Block, error) {
	value, err := s.repo.Get(block)
	if err != nil {
		return nil, err
	}
	return value, nil
}
