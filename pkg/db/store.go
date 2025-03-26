package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	pgxgeom "github.com/twpayne/pgx-geom"
	"highway-sign-portal-builder/pkg/config"
)

func NewDatabase(cfg *config.Config) (*pgx.Conn, error) {
	ctx := context.Background()

	var sslmode string
	if cfg.DBHost == "localhost" {
		sslmode = "disable"
	} else {
		sslmode = "require"
	}

	var dsn string
	if cfg.DatabaseUrl != "" {
		dsn = cfg.DatabaseUrl
	} else {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s search_path=public,sign", cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, sslmode)
	}

	db, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}

	err = pgxgeom.Register(ctx, db)

	return db, err
}

type SqlManager struct {
	*Queries
	db *pgx.Conn
}

func NewSqlManager(db *pgx.Conn) *SqlManager {
	return &SqlManager{db: db, Queries: New(db)}
}
