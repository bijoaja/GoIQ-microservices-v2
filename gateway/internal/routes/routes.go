package routes

import (
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("secret-demo-key")

func Register(r *gin.Engine) {
	r.POST("/login", proxy("http://auth-service:8082/login"))

	protected := r.Group("/")
	protected.Use(AuthMiddleware())
	protected.GET("/users", proxy("http://user-service:8081/users"))
}

func proxy(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		client := &http.Client{}
		req, err := http.NewRequest(c.Request.Method, target, c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}
		for k, v := range c.Request.Header {
			for _, vv := range v {
				req.Header.Add(k, vv)
			}
		}
		resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing auth"})
			return
		}
		parts := strings.SplitN(auth, " ", 2)
		if len(parts) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "bad auth"})
			return
		}
		_, err := jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		c.Next()
	}
}
