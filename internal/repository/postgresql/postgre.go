package repository

import (
	"fmt"
	"log"

	"github.com/AdiKhoironHasan/golangProject1/internal/models"
	"github.com/AdiKhoironHasan/golangProject1/internal/repository"

	dsnErrors "github.com/AdiKhoironHasan/golangProject1/pkg/errors"
	mhsErrors "github.com/AdiKhoironHasan/golangProject1/pkg/errors"

	"github.com/jmoiron/sqlx"
)

const (
	SaveMahasiswa          = `INSERT INTO kampus.mahasiswas (nama, nim, created_at) VALUES ($1, $2, now()) RETURNING id`
	SaveMahasiswaAlamat    = `INSERT INTO kampus.mahasiswa_alamats (jalan, no_rumah, created_at, id_mahasiswas) VALUES ($1,$2, now(), $3)`
	UpdateMahasiswaNama    = `UPDATE kampus.mahasiswas SET nama = $1, updated_at = now() where id = $2`
	SaveAlamatId           = `INSERT INTO kampus.mahasiswa_alamats (jalan, no_rumah, created_at, id_mahasiswas) VALUES ($1,$2, now(), $3)`
	ShowAllMahasiswa       = `SELECT id, nama, nim FROM kampus.mahasiswas`
	ShowAllAlamat          = `SELECT id_mahasiswas, jalan, no_rumah FROM kampus.mahasiswa_alamats`
	ShowAllMahasiswaAlamat = `SELECT a.id, a.nama, a.nim, b.jalan, b.no_rumah from kampus.mahasiswas a JOIN kampus.mahasiswa_alamats b ON a.id = b.id_mahasiswas`

	SaveDosen       = `INSERT INTO kampus.dosens (nama, nidn, created_at) VALUES ($1, $2, now()) RETURNING id`
	SaveDosenAlamat = `INSERT INTO kampus.dosen_alamats (jalan, no_rumah, created_at, id_dosens) VALUES ($1,$2, now(), $3)`
)

var statement PreparedStatement

type PreparedStatement struct {
	updateMahasiswaNama    *sqlx.Stmt //membungkus query untuk melindungi dari sql inject
	saveAlamatId           *sqlx.Stmt
	showAllMahasiswa       *sqlx.Stmt
	showAllAlamat          *sqlx.Stmt
	showAllMahasiswaAlamat *sqlx.Stmt
}

type PostgreSQLRepo struct {
	Conn *sqlx.DB
}

func NewRepo(Conn *sqlx.DB) repository.Repository {

	repo := &PostgreSQLRepo{Conn}
	InitPreparedStatement(repo)
	return repo
}

func (p *PostgreSQLRepo) Preparex(query string) *sqlx.Stmt {
	statement, err := p.Conn.Preparex(query)
	if err != nil {
		log.Fatalf("Failed to preparex query: %s. Error: %s", query, err.Error())
	}

	return statement
}

func InitPreparedStatement(m *PostgreSQLRepo) {
	statement = PreparedStatement{
		updateMahasiswaNama:    m.Preparex(UpdateMahasiswaNama),
		saveAlamatId:           m.Preparex(SaveAlamatId),
		showAllMahasiswa:       m.Preparex(ShowAllMahasiswa),
		showAllAlamat:          m.Preparex(ShowAllAlamat),
		showAllMahasiswaAlamat: m.Preparex(ShowAllMahasiswaAlamat),
	}
}

func (p *PostgreSQLRepo) SaveMahasiswaAlamat(dataMahasiswa *models.MahasiswaModels, dataAlamat []*models.MahasiswaAlamatModels) error {

	tx, err := p.Conn.Beginx()
	if err != nil {
		log.Println("Failed Begin Tx SaveMahasiswa Alamat : ", err.Error())
		return fmt.Errorf(mhsErrors.ErrorDB)
	}
	var idMahasiswa int64
	err = tx.QueryRow(SaveMahasiswa, dataMahasiswa.Name, dataMahasiswa.Nim).Scan(&idMahasiswa) //scan untuk mencari data yang di returning dari query

	if err != nil {
		tx.Rollback()
		log.Println("Failed Query SaveMahasiswa: ", err.Error())
		return fmt.Errorf(mhsErrors.ErrorDB)
	}

	for _, val := range dataAlamat {
		_, err = tx.Exec(SaveMahasiswaAlamat, val.Jalan, val.NoRumah, idMahasiswa) //exec itu ga perlu scan seperti queryrow
		if err != nil {
			tx.Rollback() //hapus semua yang sudah di query sampai setelah begin
			log.Println("Failed Query SaveMahasiswaAlamat : ", err.Error())
			return fmt.Errorf(mhsErrors.ErrorDB)
		}
	}

	return tx.Commit() //untuk patenkan atau simpan query(data) ke db
}

func (p *PostgreSQLRepo) ShowAllMahasiswaAlamat() ([]*models.ShowMahasiswaAlamatModels, error) {
	// var dataMahasiswas []*models.MahasiswaModels

	// err := statement.showAllMahasiswa.Select(&dataMahasiswas)
	// if err != nil {
	// 	log.Println("Failed Query ShowAllMahasiswa : ", err.Error())
	// 	return nil, nil, fmt.Errorf(mhsErrors.ErrorDB)
	// }
	// fmt.Println("data : ", dataMahasiswas)

	// var dataAlamat []*models.MahasiswaAlamatModels
	// err = statement.showAllAlamat.Select(&dataAlamat)
	// if err != nil {
	// 	log.Println("Failed Query ShowAllAlamat : ", err.Error())
	// 	return nil, nil, fmt.Errorf(mhsErrors.ErrorDB)
	// }
	// fmt.Println("data : ", dataAlamat)

	// Data := make([]models.Mahasiswas, len(dataMahasiswas))

	// for i, datas := range dataMahasiswas {
	// 	Data[i].id = datas.ID
	// 	Data[i].nama = datas.Name
	// 	Data[i].nim = datas.Nim
	// }

	// x := reflect.TypeOf(dataAlamat).Kind()
	// fmt.Println(x)

	var AllMahasiswaAlamat []*models.ShowMahasiswaAlamatModels

	err := statement.showAllMahasiswaAlamat.Select(&AllMahasiswaAlamat)
	if err != nil {
		log.Println("Failed Query ShowAllMahasiswaAlamat : ", err.Error())
		return nil, fmt.Errorf(mhsErrors.ErrorDB)
	}

	fmt.Println(AllMahasiswaAlamat)

	return AllMahasiswaAlamat, nil
}

func (p *PostgreSQLRepo) UpdateMahasiswaNama(dataMahasiswa *models.MahasiswaModels) error {

	result, err := statement.updateMahasiswaNama.Exec(dataMahasiswa.Name, dataMahasiswa.ID)

	if err != nil {
		log.Println("Failed Query UpdateMahasiswaNama : ", err.Error())
		return fmt.Errorf(mhsErrors.ErrorDB)
	}

	rows, err := result.RowsAffected()

	if err != nil {
		log.Println("Failed RowAffectd UpdateMahasiswaNama : ", err.Error())
		return fmt.Errorf(mhsErrors.ErrorDB)
	}

	if rows < 1 {
		log.Println("UpdateMahasiswaNama: No Data Changed")
		return fmt.Errorf(mhsErrors.ErrorNoDataChange)
	}

	return nil
}

func (p *PostgreSQLRepo) SaveAlamatId(dataAlamat *models.MahasiswaAlamatModels) error {
	result, err := statement.saveAlamatId.Exec(dataAlamat.Jalan, dataAlamat.NoRumah, dataAlamat.IDMahasiswas)

	if err != nil {
		log.Println("Failed Query SaveAlamatId : ", err.Error())
		return fmt.Errorf(mhsErrors.ErrorDB)
	}

	rows, err := result.RowsAffected()

	if err != nil {
		log.Println("Failed RowAffectd SaveAlamatId : ", err.Error())
		return fmt.Errorf(mhsErrors.ErrorDB)
	}

	if rows < 1 {
		log.Println("SaveAlamatId: No Data Changed")
		return fmt.Errorf(mhsErrors.ErrorNoDataChange)
	}

	return nil
}

// DOSEN

func (p *PostgreSQLRepo) SaveDosenAlamat(dataDosen *models.DosenModels, dataAlamat []*models.DosenAlamatModels) error {
	tx, err := p.Conn.Beginx()
	if err != nil {
		log.Println("Failed Begin Tx SaveDosen Alamat : ", err.Error())
		return fmt.Errorf(dsnErrors.ErrorDB)
	}

	var idDosen int64
	tx.QueryRow(SaveDosen, dataDosen.Name, dataDosen.Nidn).Scan(&idDosen)

	if err != nil {
		tx.Rollback()
		log.Println("Failed Query SaveDosen: ", err.Error())
		return fmt.Errorf(dsnErrors.ErrorDB)
	}

	for _, val := range dataAlamat {
		_, err = tx.Exec(SaveDosenAlamat, val.Jalan, val.NoRumah, idDosen)
		if err != nil {
			tx.Rollback()
			log.Println("Failed Query SaveDosenAlamat : ", err.Error())
			return fmt.Errorf(dsnErrors.ErrorDB)
		}
	}

	return tx.Commit()
}
