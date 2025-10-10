package routes

import (
	"github.com/example/micro/user-service/internal/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(r *gin.Engine, db *gorm.DB) {
	routes := r.Group("/users")
	routes.GET("", controllers.ListUsers(db))
	routes.POST("", controllers.CreateUser(db))
	routes.GET(":id", func(c *gin.Context) { c.JSON(200, gin.H{"msg":"not implemented"}) })
}
