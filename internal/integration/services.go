package integration

import "github.com/AdiKhoironHasan/golangProject1/pkg/dto"

type IntegServices interface {
	GetRandomDadJokes(req *dto.GetDadJokesInternalReqDTO) (*dto.GetDadJokesRandomRespDTO, error)
}
