package http

import (
	"github.com/Mohsen20031203/learn-gochain-core/config"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config  config.Config
	router  *gin.Engine
	handler *Handler
}

func NewServer(cfg config.Config, handler *Handler) *Server {
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
