package blockchain

import (
	"encoding/json"

	"github.com/Mohsen20031203/learn-gochain-core/internal/domain/transaction"
	"github.com/Mohsen20031203/learn-gochain-core/internal/infrastructure/network"
)

func (s *NodeService) BroadcastTrx(trx *[]transaction.Transaction) {
	if s.broadcaster == nil {
		return
	}

	data, err := json.Marshal(trx)
	if err != nil {
		return
	}

	msg := network.Message{
		Type: "tx",
		Data: data,
	}

	s.broadcaster.Broadcast(msg)
}
