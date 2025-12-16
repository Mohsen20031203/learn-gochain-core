package http

import (
	"net/http"

	"github.com/Mohsen20031203/learn-gochain-core/internal/domain/block"
	"github.com/Mohsen20031203/learn-gochain-core/internal/usecase/blockchain"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	node *blockchain.NodeService
}

func NewHandler(node *blockchain.NodeService) *Handler {
	return &Handler{node: node}
}

func (h *Handler) GetChain(c *gin.Context) {
	chain, err := h.node.GetChain()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"length": len(chain),
		"chain":  chain,
	})
}

func (h *Handler) CreateBlock(c *gin.Context) {
	var b block.Block

	if err := c.BindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	newBlock, err := h.node.AddBlock(b.Data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newBlock)
}
