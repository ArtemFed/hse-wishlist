package middleware

import (
	"github.com/gin-gonic/gin"
)

func Swagger(g *gin.RouterGroup) {
	//url := ginSwagger.URL("http://localhost:8082/swagger/doc.json")
	//g.GET("/swagger-ui", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	//path := "./services/tasks/.codegen/task-codegen.yaml"
	//g.GET("/swagger-ui", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL(path)))
	//fmt.Println("Path" + path)
	//if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
	//	fmt.Println("Path FILE DOES NOT EXISTS" + path)
	//}
}
