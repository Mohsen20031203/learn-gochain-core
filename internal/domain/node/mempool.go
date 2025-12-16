package node

import "github.com/Mohsen20031203/learn-gochain-core/internal/domain/transaction"

// HasTransactionMempool checks if a transaction with the given ID exists in the mempool
func (n *Node) HasTransactionMempool(txID string) bool {
	return n.mempool.HasTransaction(txID)
}

// AddTransactionMempool adds a new transaction to the mempool
func (n *Node) AddTransactionMempool(tx transaction.Transaction) {
	n.mempool.AddTransaction(tx)
}

// GetMempoolTransactions returns all transactions currently in the mempool
func (n *Node) GetMempoolTransactions() []transaction.Transaction {
	return n.mempool.GetTransactions()
}

// ClearMempool removes all transactions from the mempool
func (n *Node) ClearMempool() {
	n.mempool.Clear()
}

// SizeMempool returns the number of transactions in the mempool
func (n *Node) SizeMempool() int {
	return n.mempool.Size()
}

// RemoveTransactionMempool removes a specific transaction from the mempool
func (n *Node) RemoveTransactionMempool(tx transaction.Transaction) {
	n.mempool.RemoveTransaction(tx)
}
