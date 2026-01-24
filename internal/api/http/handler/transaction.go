package handler

import (
	"net/http"

	"github.com/Mohsen20031203/learn-gochain-core/internal/domain/transaction"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SubmitTransactions(c *gin.Context) {
	var rep []transaction.Transaction

	if err := c.BindJSON(&rep); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := h.node.SubmitTransactions(rep)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	h.node.BroadcastTrx(&rep)

	c.JSON(http.StatusOK, gin.H{"message": "transaction added to mempool"})
}
