package dto

import (
	"github.com/AdiKhoironHasan/golangProject1/pkg/common/validator"
)

type MahasiswaReqDTO struct {
	Nama    string         `json:"nama" valid:"required" validname:"nama"`
	Nim     string         `json:"nim" valid:"required" validname:"nim"`
	Alamats []AlamatReqDTO `json:"alamat" valid:"required" `
}

type AlamatReqDTO struct {
	Jalan   string `json:"jalan"`
	NoRumah string `json:"no_rumah"`
}

func (dto *MahasiswaReqDTO) Validate() error {
	v := validator.NewValidate(dto)

	return v.Validate()
}

type UpadeMahasiswaNamaReqDTO struct {
	Nama string `json:"nama" valid:"required" validname:"nama"`
	ID   int64  `json:"id" valid:"required,integer,non_zero" validname:"id"`
}

func (dto *UpadeMahasiswaNamaReqDTO) Validate() error {
	v := validator.NewValidate(dto)

	return v.Validate()
}

type AlamatIdReqDTO struct {
	Jalan        string `json:"jalan"`
	NoRumah      string `json:"no_rumah"`
	IDMahasiswas int64  `json:"mahasiswa_id" valid:"required,integer,non_zero" validname:"mahasiswa_id"`
}

func (dto *AlamatIdReqDTO) Validate() error {
	v := validator.NewValidate(dto)

	return v.Validate()
}
