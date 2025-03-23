package config

import (
	"github.com/ArtemFed/hse-wishlist/pkg/xapp"
	"github.com/ArtemFed/hse-wishlist/pkg/xconfig"
	"github.com/ArtemFed/hse-wishlist/pkg/xdb/postgres"
	"github.com/ArtemFed/hse-wishlist/pkg/xhttp"
	"github.com/ArtemFed/hse-wishlist/pkg/xlogger"
	"github.com/ArtemFed/hse-wishlist/pkg/xshutdown"
	"github.com/ArtemFed/hse-wishlist/pkg/xtracer"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	App              *xapp.Config      `mapstructure:"app"`
	Http             *xhttp.Config     `mapstructure:"http"`
	Logger           *xlogger.Config   `mapstructure:"logger"`
	Postgres         *postgres.Config  `mapstructure:"postgres"`
	GracefulShutdown *xshutdown.Config `mapstructure:"graceful_shutdown"`
	Tracer           *xtracer.Config   `mapstructure:"tracer"`
	//Metrics          *metrics.Config   `mapstructure:"metrics"`
}

func NewConfig(filePath string, appName string) (*Config, error) {
	viper.SetConfigFile(filePath)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error while reading config file: %v", err)
	}

	// Загрузка конфигурации в структуру Config
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("error while unmarshalling config file: %v", err)
	}

	// Замена значений из переменных окружения, если они заданы
	xconfig.ReplaceWithEnv(&config, appName)

	return &config, nil
}
