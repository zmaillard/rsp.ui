package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"highway-sign-portal-builder/pkg/config"
)

func NewDatabase(cfg *config.Config) (*pgx.Conn, error) {
	ctx := context.Background()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable sslmode=require search_path=public,sign", cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	db, err := pgx.Connect(ctx, dsn)

	return db, err
}

type SqlManager struct {
	*Queries
	db *pgx.Conn
}

func NewSqlManager(db *pgx.Conn) *SqlManager {
	return &SqlManager{db: db, Queries: New(db)}
}
