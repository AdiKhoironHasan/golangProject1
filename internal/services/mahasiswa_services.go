package services

import (
	"github.com/AdiKhoironHasan/golangProject1/internal/repository"
)

type service struct {
	mysqlrepo repository.Repository
}

func NewService(mysqlrepo repository.Repository) Services {
	return &service{mysqlrepo}
}
