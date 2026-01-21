package blockchain

import "github.com/Mohsen20031203/learn-gochain-core/internal/domain/transaction"

func (s *NodeService) GetMempoolTransactions() []transaction.Transaction {
	return s.node.GetMempoolTransactions()
}
