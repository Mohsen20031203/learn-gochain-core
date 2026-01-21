package blockchain

import (
	"encoding/json"
	"fmt"

	"github.com/Mohsen20031203/learn-gochain-core/internal/domain/block"
	"github.com/Mohsen20031203/learn-gochain-core/internal/domain/transaction"
	"github.com/Mohsen20031203/learn-gochain-core/internal/infrastructure/network"
)

func (s *NodeService) SetBroadcaster(b *network.TCPBroadcaster) {
	s.broadcaster = b
}

func (s *NodeService) HandleNodeMessage(msg network.Message) {
	switch msg.Type {
	case "block":
		var blc block.Block
		if err := json.Unmarshal(msg.Data, &blc); err != nil {
			fmt.Println("error unmarshall block from node message:", err)
			return
		}
		if !s.validataBlock(blc) {
			fmt.Println("received invalid block from peer")
			return
		}
		if err := s.saveBlock(&blc); err != nil {
			fmt.Println("error saving block from node message:", err)
			return
		}
		fmt.Println("block saved from peer:", blc.Hash)
	case "tx":
		var txs []transaction.Transaction
		if err := json.Unmarshal(msg.Data, &txs); err != nil {
			fmt.Println("error unmarshall txs from node message:", err)
			return
		}
		s.SubmitTransactions(txs)
	}
}
