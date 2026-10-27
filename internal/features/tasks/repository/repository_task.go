package tasks_repository

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RepositoryTask struct {
	db *pgxpool.Pool
}

func NewRepositoryTask() *RepositoryTask {
	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	return &RepositoryTask{
		db: pool,
	}
}
