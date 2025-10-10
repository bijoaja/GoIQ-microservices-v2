package app

import (
	"github.com/gin-gonic/gin"

	"github.com/example/micro/gateway/internal/routes"
)

type Server struct {
	R *gin.Engine
}

func New() (*Server, error) {
	r := gin.Default()
	routes.Register(r)
	return &Server{R: r}, nil
}

func (s *Server) Run() {
	s.R.Run(":8080")
}
