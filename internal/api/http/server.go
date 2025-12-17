package http

import (
	"github.com/Mohsen20031203/learn-gochain-core/config"
	"github.com/Mohsen20031203/learn-gochain-core/internal/api/http/handler"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config  config.Config
	router  *gin.Engine
	handler *handler.Handler
}

func NewServer(cfg config.Config, handler *handler.Handler) *Server {
	router := NewRouter(handler)

	return &Server{
		config:  cfg,
		router:  router,
		handler: handler,
	}
}

func (s *Server) Start() error {
	return s.router.Run(":" + s.config.Port)
}
