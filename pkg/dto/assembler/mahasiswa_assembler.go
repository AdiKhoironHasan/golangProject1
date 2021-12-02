package assembler

import (
	"github.com/AdiKhoironHasan/golangProject1/internal/models"
	"github.com/AdiKhoironHasan/golangProject1/pkg/dto"
)

func ToSaveMahasiswa(d *dto.MahasiswaReqDTO) *models.MahasiswaModels {
	return &models.MahasiswaModels{
		Name: d.Nama,
		Nim:  d.Nim,
	}
}

func ToSaveMahasiswaAlamat(d *dto.AlamatReqDTO) *models.MahasiswaAlamatModels {
	return &models.MahasiswaAlamatModels{
		Jalan:   d.Jalan,
		NoRumah: d.NoRumah,
	}
}

func ToSaveMahasiswaAlamats(datas []dto.AlamatReqDTO) []*models.MahasiswaAlamatModels {
	var mds []*models.MahasiswaAlamatModels
	for _, m := range datas {
		mds = append(mds, ToSaveMahasiswaAlamat(&m))
	}
	return mds
}

func ToUpdateMahasiswaNama(d *dto.UpadeMahasiswaNamaReqDTO) *models.MahasiswaModels {
	return &models.MahasiswaModels{
		Name: d.Nama,
		ID:   d.ID,
	}
}

func ToSaveAlamatId(d *dto.AlamatIdReqDTO) *models.MahasiswaAlamatModels {
	return &models.MahasiswaAlamatModels{
		Jalan:        d.Jalan,
		NoRumah:      d.NoRumah,
		IDMahasiswas: d.IDMahasiswas,
	}
}

// DOSEN
func ToSaveDosenAlamat(dataAlamat *dto.AlamatDosenReqDTO) *models.DosenAlamatModels {
	return &models.DosenAlamatModels{
		Jalan:   dataAlamat.Jalan,
		NoRumah: dataAlamat.NoRumah,
	}
}

func ToSaveDosenAlamats(dataAlamats []dto.AlamatDosenReqDTO) []*models.DosenAlamatModels {
	var mds []*models.DosenAlamatModels
	for _, data := range dataAlamats {
		mds = append(mds, ToSaveDosenAlamat(&data))
	}
	return mds
}

func ToSaveDosen(dataDosen *dto.DosenReqDTO) *models.DosenModels {
	return &models.DosenModels{
		Name: dataDosen.Nama,
		Nidn: dataDosen.Nidn,
	}
}
