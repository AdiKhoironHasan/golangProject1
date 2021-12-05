package dto

type ResponseDTO struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type MahasiswaAlamatResDTO struct {
	ID      int64           `json:"id_mahasiswa"`
	Nama    string          `json:"nama"`
	Nim     string          `json:"nim"`
	Alamats []*AlamatResDTO `json:"alamat"`
}

type AlamatResDTO struct {
	Jalan   string `json:"jalan"`
	NoRumah string `json:"no_rumah"`
}

// DOSEN
type DosenAlamatResDTO struct {
	ID      int64                `json:"id_dosen"`
	Nama    string               `json:"nama"`
	Nidn    string               `json:"nidn"`
	Alamats []*AlamatDosenResDTO `json:"alamats"`
}

type AlamatDosenResDTO struct {
	Jalan   string `json:"jalan"`
	NoRumah string `json:"no_rumah"`
}
