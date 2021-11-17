package repository

import (
	"github.com/AdiKhoironHasan/golangProject1/internal/repository"

	"github.com/jmoiron/sqlx"
)

const ()

var statement PreparedStatement

type PreparedStatement struct {
}

type PostgreSQLRepo struct {
	Conn *sqlx.DB
}

func NewRepo(Conn *sqlx.DB) repository.Repository {

	repo := &PostgreSQLRepo{Conn}
	// InitPreparedStatement(repo)
	return repo
}
