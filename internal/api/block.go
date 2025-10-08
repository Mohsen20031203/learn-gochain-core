package api

import (
	"time"

	"github.com/Mohsen20031203/learn-gochain-core/internal/model"
	"github.com/gin-gonic/gin"
)

func (s *Server) CreatBlock(c *gin.Context) {

	var block model.Block
	err := c.BindJSON(&block)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	block.Timestamp = time.Now()

	if len(s.Chain) == 0 {
		block.Index = 0
		block.PrevHash = "0"
	} else {
		block.Index = int64(len(s.Chain))
		block.PrevHash = s.Chain[len(s.Chain)-1].Hash
	}

	block.Mine(2)

	s.Chain = append(s.Chain, block)

	c.JSON(200, gin.H{
		"message": "Block created successfully",
		"block":   block,
	})

}

func (s *Server) GetChain(c *gin.Context) {

}
