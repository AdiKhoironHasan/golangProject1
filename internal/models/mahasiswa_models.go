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

type ShowMahasiswaAlamatModels struct {
	ID      int64  `db:"id"`
	Name    string `db:"nama"`
	Nim     string `db:"nim"`
	Jalan   string `db:"jalan"`
	NoRumah string `db:"no_rumah"`
}

type GetMahasiswaAlamatsModels struct {
	ID      int64  `db:"id"`
	Name    string `db:"nama"`
	Nim     string `db:"nim"`
	Jalan   string `db:"jalan"`
	NoRumah string `db:"no_rumah"`
}

// type DBAlamat struct {
// 	id      int64
// 	jalan   string
// 	norumah string
// }

// type DBMahasiswas struct {
// 	id   int64
// 	nama string
// 	nim  string
// }

// type Alamat struct {
// 	id      int64
// 	jalan   string
// 	norumah string
// }

// type Mahasiswas struct {
// 	id      int64
// 	nama    string
// 	nim     string
// 	alamats []Alamat
// }

// type DBmhs DBMahasiswas
// type DBalmt Alamat

// DOSEN
type DosenModels struct {
	ID   int64  `db:"id"`
	Name string `db:"nama"`
	Nidn string `db:"nim"`
}

type DosenAlamatModels struct {
	ID      int64  `db:"id"`
	Jalan   string `db:"jalan"`
	NoRumah string `db:"no_rumah"`
	IdDosen int64  `db:"id_mahasiswas"` //sama dengan field db
}

type ShowAllDosenAlamatModels struct {
	ID      int64  `db:"id"`
	Name    string `db:"nama"`
	Nidn    string `db:"nidn"`
	Jalan   string `db:"jalan"`
	NoRumah string `db:"no_rumah"`
}
