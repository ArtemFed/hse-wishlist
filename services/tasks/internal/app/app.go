package app

import (
	"context"
	"fmt"
	"github.com/ArtemFed/hse-wishlist/pkg/xdb/postgres"
	"github.com/ArtemFed/hse-wishlist/pkg/xerror"
	"github.com/ArtemFed/hse-wishlist/pkg/xshutdown"
	"github.com/ArtemFed/hse-wishlist/pkg/xtracer"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/config"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/handler/http"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/log"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/repository/postgre"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/service"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"
)

type App struct {
	cfg            *config.Config
	handler        http.Handler
	address        string
	tracerProvider *trace.TracerProvider
}

func NewApp(cfg *config.Config) (*App, error) {

	// INFRASTRUCTURE ----------------------------------------------------------------------

	err := log.Init(cfg.Logger, cfg.App)
	if err != nil {
		return nil, errors.Wrap(err, "Init Logger")
	}
	// Чистим кэш logger при shutdown
	xshutdown.AddCallback(
		&xshutdown.Callback{
			Name: "ZapLoggerCacheWipe",
			FnCtx: func(ctx context.Context) error {
				return log.Logger.Sync()
			},
		})
	log.Logger.Info("Init Logger – success")

	// Инициализируем обработку ошибок
	err = xerror.InitAppError(cfg.App)
	if err != nil {
		log.Logger.Fatal("while initializing App Error handling package", zap.Error(err))
	}

	// Инициализируем трассировку
	var tp *trace.TracerProvider = nil
	if cfg.Tracer.Enable {
		tp, err = xtracer.Init(cfg.Tracer, cfg.App)
		if err != nil {
			return nil, err
		}
		xshutdown.AddCallback(
			&xshutdown.Callback{
				Name: "OpenTelemetryShutdown",
				FnCtx: func(ctx context.Context) error {
					if err := tp.Shutdown(context.Background()); err != nil {
						log.Logger.Error("Error shutting down tracer provider: %v", zap.Error(err))
						return err
					}
					return nil
				},
			})
		log.Logger.Info("Init Tracer – success")
	} else {
		log.Logger.Info("Init Tracer – skipped")
	}

	// Инициализируем Prometheus
	//if cfg.Metrics.Enable {
	//	metrics.InitOnce(cfg.Metrics, log.Logger, metrics.AppInfo{
	//		Name:        cfg.App.Name,
	//		Environment: string(cfg.App.Environment),
	//		Version:     cfg.App.Version,
	//	})
	//	log.Logger.Info("Init Metrics – success")
	//} else {
	//	log.Logger.Info("Init Metrics – skipped")
	//}

	// REPOSITORY ----------------------------------------------------------------------

	// Инициализация PostgreSQL
	postgresDb, err := postgres.NewDB(cfg.Postgres)
	if err != nil {
		log.Logger.Fatal("Error init Postgres DB:", zap.Error(err))
		return nil, errors.Wrap(err, "Init Postgres DB")
	}

	// Инициализация всех репозиториев
	taskRepo := postgre.NewTaskRepository(postgresDb)
	accountRepo := postgre.NewAccountRepository(postgresDb)

	// SERVICE LAYER ----------------------------------------------------------------------

	// Name layer
	taskService := service.NewTaskService(taskRepo)
	accountService := service.NewAccountService(accountRepo)
	authService := service.NewAuthService(accountService)

	log.Logger.Info(fmt.Sprintf("Init %s – success", cfg.App.Name))

	// TRANSPORT LAYER ----------------------------------------------------------------------

	mainHandler := http.NewHandler(
		cfg,
		taskService,
		accountService,
		authService,
	)

	// инициализируем адрес сервера
	address := fmt.Sprintf(":%s", cfg.Http.Port)

	return &App{
		cfg:            cfg,
		handler:        mainHandler,
		address:        address,
		tracerProvider: tp,
	}, nil
}
