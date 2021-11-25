package services

import "github.com/AdiKhoironHasan/golangProject1/pkg/dto"

type Services interface {
	SaveMahasiswaAlamat(req *dto.MahasiswaReqDTO) error
	UpdateMahasiswaNama(req *dto.UpadeMahasiswaNamaReqDTO) error
}
