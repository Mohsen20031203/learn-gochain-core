package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Mohsen20031203/learn-gochain-core/internal/usecase/blockchain"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	node *blockchain.NodeService
}

func NewHandler(node *blockchain.NodeService) *Handler {
	return &Handler{node: node}
}

func (h *Handler) ShowMessage(c *gin.Context) {
	c.JSON(200, gin.H{"messages": h.node.Message})
}
func (h *Handler) SendMassage(c *gin.Context) {
	var msg blockchain.Message

	if err := c.BindJSON(&msg); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	if msg.From == h.node.GetConfig().NodeID {
		c.JSON(200, gin.H{"status": "ignored self"})
		return
	}

	if h.node.SeenMessages[msg.ID] {
		c.JSON(200, gin.H{"status": "duplicate"})
		return
	}

	h.node.SeenMessages[msg.ID] = true
	h.node.Message = append(h.node.Message, msg)

	for _, peer := range h.node.GetConfig().Peers {
		go func(p string) {
			url := fmt.Sprintf("http://%s/peers", p)
			body, _ := json.Marshal(msg)
			http.Post(url, "application/json", bytes.NewBuffer(body))
		}(peer)
	}

	c.JSON(200, gin.H{"status": "broadcasted"})
}
