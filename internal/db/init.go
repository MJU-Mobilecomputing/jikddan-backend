package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

var BASE_URL = "postgres://%s:%s@%s:%s/%s?sslmode=disable"
var (
	db   *pgxpool.Pool
	once sync.Once
)

func InitDB(ctx context.Context, dbConfig config.DB) (*pgxpool.Pool, error) {
	if db == nil {
		once.Do(func() {
			conn, err := setDBByConfig(ctx, dbConfig)
			if err != nil {
				log.Fatal(err)
			}
			db = conn
		})
	} else {
		return nil, errors.New("db already configured")
	}
	return db, nil
}

func setDBByConfig(ctx context.Context, c config.DB) (*pgxpool.Pool, error) {
	dbURL := fmt.Sprintf(BASE_URL, c.Username, c.Password, c.Host, c.Port, c.Name)
	conn, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return nil, err
	}
	return conn, err
}
