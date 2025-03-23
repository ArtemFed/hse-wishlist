package app

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/ArtemFed/hse-wishlist/pkg/xhttp"
	"github.com/ArtemFed/hse-wishlist/pkg/xshutdown"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/handler/http"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/log"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	requestid "github.com/sumit-tembe/gin-requestid"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"io/ioutil"
	nhttp "net/http"
	"os"
	"sigs.k8s.io/yaml"
	"time"
)

// Start - Единая точка запуска приложения
func (a *App) Start(ctx context.Context) {

	go a.startHTTPServer(ctx)

	if err := xshutdown.Wait(a.cfg.GracefulShutdown); err != nil {
		log.Logger.Error(fmt.Sprintf("Failed to gracefully shutdown %s app: %s", a.cfg.App.Name, err.Error()))
	} else {
		log.Logger.Info("App gracefully stopped")
	}
}

func (a *App) startHTTPServer(ctx context.Context) {
	// Создаем общий роутинг http сервера
	router := xhttp.NewRouter()

	// Добавление маршрута для спецификации OpenAPI в формате JSON
	router.Router().GET("/api/v1/swagger/swagger.json", func(c *gin.Context) {
		// Определите путь к вашему файлу спецификации OpenAPI
		yamlFilePath := "./services/tasks/.codegen/task-codegen.yaml"
		if os.Getenv("SWAGGER_FILE_PATH") != "" {
			yamlFilePath = os.Getenv("SWAGGER_FILE_PATH")
		}

		// Читаем содержимое файла спецификации
		yamlContent, err := ioutil.ReadFile(yamlFilePath)
		if err != nil {
			c.JSON(nhttp.StatusInternalServerError, gin.H{"error": "Failed to read OpenAPI specification file"})
			return
		}

		// Преобразуем YAML в JSON
		jsonContent, err := yaml.YAMLToJSON(yamlContent)
		if err != nil {
			c.JSON(nhttp.StatusInternalServerError, gin.H{"error": "Failed to convert YAML to JSON"})
			return
		}

		// Отправляем содержимое файла как ответ
		c.Data(nhttp.StatusOK, "application/json", jsonContent)
	})

	// Добавление маршрутов для Swagger UI
	router.Router().GET("/api/v1/swagger-ui/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/api/v1/swagger/swagger.json")))

	//router.Router().GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	var middlewares []http.MiddlewareFunc
	if a.cfg.Tracer.Enable {
		middlewares = append(middlewares, http.MiddlewareFunc(otelgin.Middleware(a.cfg.App.Name, otelgin.WithTracerProvider(a.tracerProvider))))
	}
	middlewares = append(middlewares, http.MiddlewareFunc(ginzap.Ginzap(log.Logger, time.RFC3339, true)))
	middlewares = append(middlewares, http.MiddlewareFunc(requestid.RequestID(nil)))

	http.InitHandler(a.handler, router.Router(), middlewares, "hse")

	srv := xhttp.NewServer(a.cfg.Http, router)

	// Стартуем
	log.Logger.Info(fmt.Sprintf("Starting %s HTTP server at %s:%s", a.cfg.App.Name, a.cfg.Http.Host, a.cfg.Http.Port))
	if err := srv.Start(); err != nil {
		log.Logger.Error(fmt.Sprintf("Fail with %s HTTP server: %s", a.cfg.App.Name, err.Error()))
		xshutdown.Now()
	}
}
