package blockchain

import (
	"github.com/Mohsen20031203/learn-gochain-core/internal/domain/block"
	"github.com/Mohsen20031203/learn-gochain-core/internal/domain/transaction"
)

type Node interface {
	AddTransactionMempool(tx transaction.Transaction)
	GetMempoolTransactions() []transaction.Transaction
	ClearMempool()
	SizeMempool() int
	RemoveTransactionMempool(tx transaction.Transaction)
	GetMempoolTransaction(count int) []transaction.Transaction
	IsValidNewBlockChain(blockHash block.Block) bool
	UpdateChain(blockHash block.Block)
	GetChainLastBlockHash() string
	CountBlocksinChain() int
	GetChainDifficulty() int
	MineBlock(b *block.Block)
	IsValidPoW(b *block.Block) bool
}
