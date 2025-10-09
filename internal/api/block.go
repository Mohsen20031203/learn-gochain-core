package api

import (
	"time"

	"github.com/Mohsen20031203/learn-gochain-core/internal/block"
	"github.com/gin-gonic/gin"
)

func (s *Server) CreatBlock(c *gin.Context) {

	var block block.Block
	err := c.BindJSON(&block)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	block.Timestamp = time.Now()

	//TODO: find hash from black
	//TODO: chech block with node and rolus
	//TODO: get trasaction from pool
	if len(s.Chain) == 0 {
		block.Index = 0
		block.PrevHash = "0"
	} else {
		block.Index = int64(len(s.Chain))
		block.PrevHash = s.Chain[len(s.Chain)-1].Hash
	}

	block.Mine(6)

	s.Chain = append(s.Chain, block)

	c.JSON(200, gin.H{
		"message": "Block created successfully",
		"block":   block,
	})

}

func (s *Server) GetChain(c *gin.Context) {

	c.JSON(200, gin.H{
		"length": len(s.Chain),
		"chain":  s.Chain,
	})
}
