package mempool

import (
	"github.com/Mohsen20031203/learn-gochain-core/internal/domain/transaction"
)

type MempoolDB interface {
	AddTransaction(tx transaction.Transaction)
	GetTransactions() []transaction.Transaction
	HasTransaction(txID string) bool
	Size() int
	RemoveTransaction(tx transaction.Transaction)
	Clear()
	GetTransactionsCount(count int) []transaction.Transaction
}
