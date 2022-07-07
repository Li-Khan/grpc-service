package postgres

import (
	"context"
	"fmt"
	"github.com/Li-Khan/grpc-service/configs"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

func NewPostgresRepository(config *configs.Config) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", config.UserDB, config.PasswordDB, config.HostDB, config.PortDB, config.NameDB)

	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}

	err = db.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
