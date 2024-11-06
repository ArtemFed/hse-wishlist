package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "log"
)

type Config struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Name     string `mapstructure:"name"`
	Password string `mapstructure:"password"`
	SSL      string `mapstructure:"ssl"`
}

func NewDB(cfg *Config) (*sqlx.DB, error) {
	dbDSN := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Name,
		cfg.SSL,
		cfg.Password,
	)
	log.Printf("DB DSN: '%s'\n", dbDSN)
	db, err := sqlx.Open("postgres", dbDSN)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}
