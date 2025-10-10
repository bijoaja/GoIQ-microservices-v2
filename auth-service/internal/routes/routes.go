package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var jwtKey = []byte("secret-demo-key")

func Register(r *gin.Engine) {
	r.POST("/login", func(c *gin.Context) {
		var in struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&in); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if in.Password != "password" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": in.Username,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})
		ts, err := token.SignedString(jwtKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot create token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": ts})
	})
}
