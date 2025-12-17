package handler

import (
	"github.com/Mohsen20031203/learn-gochain-core/internal/usecase/blockchain"
)

type Handler struct {
	node *blockchain.NodeService
}

func NewHandler(node *blockchain.NodeService) *Handler {
	return &Handler{node: node}
}
