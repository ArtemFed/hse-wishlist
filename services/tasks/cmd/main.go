package main

import (
	"fmt"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/app"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/config"
	"github.com/joho/godotenv"
	"golang.org/x/net/context"
	"log"
	"os"
)

const MainEnvName = ".env"
const ServiceName = "Tasks"
const ServiceCapsName = "TASKS"

func init() {
	envPath := fmt.Sprintf("services/tasks/%s", MainEnvName)
	if err := godotenv.Load(envPath); err != nil {
		log.Print(fmt.Sprintf("No '%s' file found", MainEnvName))
	}
}

func main() {
	FAIL

	ctx := context.Background()

	cfgEnvName := "CONFIG_" + ServiceCapsName
	configPath := os.Getenv(cfgEnvName)
	log.Printf("%s config path (%s): %s", ServiceName, cfgEnvName, configPath)

	// Собираем конфиг приложения
	cfg, err := config.NewConfig(configPath, ServiceCapsName)
	if err != nil {
		log.Fatalf("Fail to parse %s config: %v", ServiceName, err)
	}

	// Создаем наше приложение
	application, err := app.NewApp(cfg)
	if err != nil {
		log.Fatal(fmt.Sprintf("Fail to create '%s' service: %s", cfg.App.Name, err))
	}

	// Запускаем приложение
	application.Start(ctx)
}
