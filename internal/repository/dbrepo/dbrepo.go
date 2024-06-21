package dbrepo

import (
	"github.com/jackc/pgx/v5"
	"github.com/otavio-Pucharelli/filhos-da-luz/internal/config"
	"github.com/otavio-Pucharelli/filhos-da-luz/internal/repository"
)

// PostgresDBRepo is a repository type
type PostgresDBRepo struct {
	App *config.AppConfig
	DB  *pgx.Conn
}

// NewPostgresRepo creates a new repository
func NewPostgresRepo(conn *pgx.Conn, a *config.AppConfig) repository.DatabaseRepo {
	return &PostgresDBRepo{
		App: a,
		DB:  conn,
	}
}
