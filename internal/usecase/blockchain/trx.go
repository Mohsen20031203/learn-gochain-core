package blockchain

import "github.com/Mohsen20031203/learn-gochain-core/internal/domain/transaction"

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
