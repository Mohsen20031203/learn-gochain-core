package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetMempool(c *gin.Context) {
	mempool := h.node.GetMempoolTransactions()
	c.JSON(http.StatusOK, gin.H{
		"length":       len(mempool),
		"mempool_txns": mempool,
	})
}
