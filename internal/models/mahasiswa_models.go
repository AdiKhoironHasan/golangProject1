package models

type MahasiswaModels struct {
	ID   int64  `db:"id"`
	Name string `db:"nama"`
	Nim  string `db:"nim"`
}

type MahasiswaAlamatModels struct {
	ID           int64  `db:"id"`
	Jalan        string `db:"jalan"`
	NoRumah      string `db:"no_rumah"`
	IDMahasiswas int64  `db:"id_mahasiswas"` //sama dengan field db
}

type DBAlamat struct {
	id      int64
	jalan   string
	norumah string
}

type DBMahasiswas struct {
	id   int64
	nama string
	nim  string
}

type Alamat struct {
	id      int64
	jalan   string
	norumah string
}

type Mahasiswas struct {
	id      int64
	nama    string
	nim     string
	alamats []Alamat
}

type DBmhs DBMahasiswas
type DBalmt Alamat
