package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Swagger(g *gin.RouterGroup) {
	//url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	g.GET("/swagger-ui", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
