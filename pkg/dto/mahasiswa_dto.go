package dto

import (
	"github.com/AdiKhoironHasan/golangProject1/pkg/common/validator"
)

type MahasiswaReqDTO struct {
	Nama    string         `json:"nama" valid:"required" validname:"nama"`
	Nim     string         `json:"nim" valid:"required" validname:"nim"`
	Alamats []AlamatReqDTO `json:"alamat" valid:"required" `
}

func (dto *MahasiswaReqDTO) Validate() error {
	v := validator.NewValidate(dto)

	return v.Validate()
}

type AlamatReqDTO struct {
	Jalan   string `json:"jalan"`
	NoRumah string `json:"no_rumah"`
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

// DOSEN
type DosenReqDTO struct {
	Nama    string              `json:"nama" valid:"required" validname:"nama"`
	Nidn    string              `json:"nidn" valid:"required" validname:"nidn"`
	Alamats []AlamatDosenReqDTO `json:"alamat" valid:"required" `
}

func (dto *DosenReqDTO) Validate() error {
	v := validator.NewValidate(dto)

	return v.Validate()
}

type AlamatDosenReqDTO struct {
	Jalan   string `json:"jalan"`
	NoRumah string `json:"no_rumah"`
}

type UpdateDosenNamaReqDTO struct {
	Nama string `json:"nama" valid:"required" validname:"nama"`
	ID   int64  `json:"id" valid:"required,integer,non_zero" validname:"id"`
}

func (dto *UpdateDosenNamaReqDTO) Validate() error {
	v := validator.NewValidate(dto)

	return v.Validate()
}

type AlamatDosenByIDReqDTO struct {
	Jalan   string `json:"jalan"`
	NoRumah string `json:"no_rumah"`
	IdDosen int64  `json:"id_dosen" valid:"required,integer,non_zero" validname:"id_dosen"`
}

func (dto *AlamatDosenByIDReqDTO) Validate() error { //method yang menempel ke AlamatDosen
	v := validator.NewValidate(dto)

	return v.Validate()
}

type DosenParamReqDTO struct {
	IdDosen int64  `json:"id_dosen" validname:"id_dosen" query:"id_dosen"`
	Nama    string `json:"nama" validname:"nama" query:"nama"`
	Nidn    string `json:"nidn" validname:"nidn" query:"nidn"`
}

func (dto *DosenParamReqDTO) Validate() error {
	v := validator.NewValidate(dto)

	return v.Validate()
}

// dosen by param baiknya digabung dosenreq apa buat baru
// fungsi validname required, nonzero
