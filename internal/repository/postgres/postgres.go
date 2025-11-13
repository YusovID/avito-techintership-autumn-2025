package postgres

import (
	"fmt"
	"log/slog"

	"github.com/YusovID/pr-reviewer-service/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDB(cfg config.Postgres, log *slog.Logger) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database,
	)

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("can't connect to database: %w", err)
	}

	return db, nil
}
