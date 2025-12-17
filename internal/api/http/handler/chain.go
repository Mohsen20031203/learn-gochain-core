package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
