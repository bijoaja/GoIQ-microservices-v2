package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/bijoaja/GoIQ-microservices.v2/user-service/internal/database"
	"github.com/bijoaja/GoIQ-microservices.v2/user-service/internal/routes"
)

type Server struct {
	DB *gorm.DB
	R  *gin.Engine
}

func New() (*Server, error) {
	db := database.Connect()
	r := gin.Default()
	routes.Register(r, db)
	return &Server{DB: db, R: r}, nil
}

func (s *Server) Run() {
	fmt.Println("user-service running on :8081")
	s.R.Run(":8081")
}

func (s *Server) Migrate() error {
	return database.Migrate(s.DB)
}
