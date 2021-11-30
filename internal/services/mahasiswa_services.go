package services

import (
	"github.com/AdiKhoironHasan/golangProject1/internal/repository"
	"github.com/AdiKhoironHasan/golangProject1/pkg/dto"
	"github.com/AdiKhoironHasan/golangProject1/pkg/dto/assembler"
)

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Services {
	return &service{repo}
}

func (s *service) SaveMahasiswaAlamat(req *dto.MahasiswaReqDTO) error {

	dtAlamat := assembler.ToSaveMahasiswaAlamats(req.Alamats)
	dtMahasiswa := assembler.ToSaveMahasiswa(req)

	err := s.repo.SaveMahasiswaAlamat(dtMahasiswa, dtAlamat)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) ShowAllMahasiswaAlamat() (string, error) {

	data, err := s.repo.ShowAllMahasiswaAlamat()

	return data, err
}

func (s *service) UpdateMahasiswaNama(req *dto.UpadeMahasiswaNamaReqDTO) error {

	dtMhsiswa := assembler.ToUpdateMahasiswaNama(req)

	err := s.repo.UpdateMahasiswaNama(dtMhsiswa)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) SaveAlamatId(req *dto.AlamatIdReqDTO) error {
	dtAlamat := assembler.ToSaveAlamatId(req)

	err := s.repo.SaveAlamatId(dtAlamat)
	if err != nil {
		return err
	}

	return nil
}
