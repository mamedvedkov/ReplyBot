package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Config struct {
	Host     string `default:"192.168.1.104" split_words:"true"`
	Port     uint   `default:"5432" split_words:"true"`
	User     string `default:"postgres" split_words:"true"`
	Password string `default:"postgres" split_words:"true"`
	Database string `default:"postgres" split_words:"true"`
	MaxConns uint   `default:"10" split_words:"true"`
}

type PG struct {
	pool *pgxpool.Pool
}

func Must(cfg Config) *PG {
	pg, err := New(cfg)
	if err != nil {
		panic(err)
	}

	return pg
}

func New(cfg Config) (*PG, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf("postgres://%s:%s@%s:%v/%s?sslmode=disable&pool_max_conns=%v",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.MaxConns))
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.ConnectConfig(context.TODO(), config)
	if err != nil {
		return nil, err
	}

	return &PG{pool: pool}, nil
}
