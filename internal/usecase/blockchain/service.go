package blockchain

import (
	"github.com/Mohsen20031203/learn-gochain-core/config"
	"github.com/Mohsen20031203/learn-gochain-core/internal/domain/node"
	"github.com/Mohsen20031203/learn-gochain-core/internal/infrastructure/network"
	"github.com/Mohsen20031203/learn-gochain-core/internal/infrastructure/storage/lvldb"
)

type NodeService struct {
	node        *node.Node
	repo        Repository
	config      config.Config
	mineTrigger chan struct{}
	broadcaster *network.TCPBroadcaster
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
