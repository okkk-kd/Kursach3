package postgres

import (
	"context"
	"dbKursach/internal/config"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/stdlib"

	"github.com/jmoiron/sqlx"
)

func InitPsqlDB(ctx context.Context, cfg *config.Config) (*sqlx.DB, error) {
	connectionURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DBName,
		cfg.Postgres.SSLMode,
	)

	database, err := sqlx.ConnectContext(ctx, cfg.Postgres.PGDriver, connectionURL)
	if err != nil {
		return nil, err
	}

	database.SetMaxOpenConns(cfg.Postgres.Settings.MaxOpenConns)
	database.SetConnMaxLifetime(cfg.Postgres.Settings.ConnMaxLifetime * time.Second)
	database.SetMaxIdleConns(cfg.Postgres.Settings.MaxIdleConns)
	database.SetConnMaxIdleTime(cfg.Postgres.Settings.ConnMaxIdleTime * time.Second)

	if err = database.Ping(); err != nil {
		return nil, err
	}

	return database, nil
}
