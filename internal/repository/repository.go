package repository

import "github.com/AdiKhoironHasan/golangProject1/internal/models"

type Repository interface {
	SaveMahasiswaAlamat(dataMahasiswa *models.MahasiswaModels, dataAlamat []*models.MahasiswaAlamatModels) error
	UpdateMahasiswaNama(dataMahasiswa *models.MahasiswaModels) error
	SaveAlamatId(dataAlamat *models.MahasiswaAlamatModels) error
	ShowAllMahasiswaAlamat() ([]*models.ShowMahasiswaAlamatModels, error)
	GetMahasiswaAlamat(where string) ([]*models.GetMahasiswaAlamatsModels, error)

	SaveDosenAlamat(dataDosen *models.DosenModels, dataAlamat []*models.DosenAlamatModels) error
	UpdateDosenNama(dataDosen *models.DosenModels) error
	SaveDosenAlamatByID(dataDosenAlamat *models.DosenAlamatModels) error
	ShowAllDosenAlamat(where string) ([]*models.ShowAllDosenAlamatModels, error)
}
