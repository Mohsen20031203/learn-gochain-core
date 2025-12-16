package node

import (
	"github.com/Mohsen20031203/learn-gochain-core/internal/domain/blockchain"
	"github.com/Mohsen20031203/learn-gochain-core/internal/domain/mempool"
)

type Node struct {
	id      string
	chain   *blockchain.Blockchain
	mempool *mempool.Mempool
}

// NewNode creates and returns a new Node instance.
// id is the unique identifier for the node ,
// difficulty is the mining difficulty for the node's blockchain
func NewNode(id string, difficulty int) *Node {
	return &Node{
		chain:   blockchain.New(difficulty),
		id:      id,
		mempool: mempool.NewMempool(),
	}
}

// GetID returns the ID of the node
func (n *Node) GetID() string {
	return n.id
}
