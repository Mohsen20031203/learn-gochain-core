package api

import (
	"github.com/Mohsen20031203/learn-gochain-core/config"
	"github.com/Mohsen20031203/learn-gochain-core/internal/model"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Config config.Config
	router *gin.Engine
	Chain  []model.Block
}

func NewServer(config config.Config) *Server {
	server := Server{Config: config}

	err := server.setupRouter()
	if err != nil {
		panic(err)
	}
	return &server
}

func (s *Server) setupRouter() error {

	router := gin.Default()
	router.Use(gin.Recovery())

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://65.109.219.253:3000", "http://65.109.219.253:3001", "http://localhost:3000", "http://localhost:3001", "https://doctor.actelmon.ir", "https://admin.actelmon.ir", "https://app.actelmon.ir", "http://192.168.254.23:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "X-Requested-With"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60,
	}))

	router.Use(gzip.Gzip(gzip.DefaultCompression))

	router.GET("/chain", s.GetChain)
	router.POST("/block", s.CreatBlock)
	s.router = router
	return nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
