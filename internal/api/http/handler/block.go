package handler

import "github.com/gin-gonic/gin"

func (h *Handler) GetBlockByHash(c *gin.Context) {
	blockID := c.Param("hash")
	block, err := h.node.GetBlockByHash(blockID)
	if err != nil {
		c.JSON(404, gin.H{"error": "block not found"})
		return
	}
	c.JSON(200, block)
}
