package services

import (
	"fmt"

	integ "github.com/AdiKhoironHasan/golangProject1/internal/integration"
	"github.com/AdiKhoironHasan/golangProject1/internal/repository"
	"github.com/AdiKhoironHasan/golangProject1/pkg/dto"
	"github.com/AdiKhoironHasan/golangProject1/pkg/dto/assembler"
)

type service struct {
	repo      repository.Repository
	IntegServ integ.IntegServices
}

func NewService(repo repository.Repository, IntegServ integ.IntegServices) Services {
	return &service{repo, IntegServ}
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

func (s *service) ShowAllMahasiswaAlamat() ([]*dto.MahasiswaAlamatResDTO, error) {

	getMahasiswaMap := make(map[int64]*dto.MahasiswaAlamatResDTO)
	DataMahasiswaAlamat, err := s.repo.ShowAllMahasiswaAlamat()
	if err != nil {
		return nil, err
	}

	for _, val := range DataMahasiswaAlamat {
		if _, ok := getMahasiswaMap[val.ID]; !ok {
			getMahasiswaMap[val.ID] = &dto.MahasiswaAlamatResDTO{
				ID:   val.ID,
				Nama: val.Name,
				Nim:  val.Nim,
			}
			getMahasiswaMap[val.ID].Alamats = append(getMahasiswaMap[val.ID].Alamats, &dto.AlamatResDTO{
				Jalan:   val.Jalan,
				NoRumah: val.NoRumah,
			})
		} else {
			getMahasiswaMap[val.ID].Alamats = append(getMahasiswaMap[val.ID].Alamats, &dto.AlamatResDTO{
				Jalan:   val.Jalan,
				NoRumah: val.NoRumah,
			})
		}

	}

	var Data []*dto.MahasiswaAlamatResDTO
	for _, datas := range getMahasiswaMap {
		Data = append(Data, datas)
	}

	return Data, err
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

// DOSEN
func (s *service) SaveDosenAlamat(req *dto.DosenReqDTO) error {
	dtaAlamat := assembler.ToSaveDosenAlamats(req.Alamats)
	dtDosen := assembler.ToSaveDosen(req)

	err := s.repo.SaveDosenAlamat(dtDosen, dtaAlamat)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateDosenNama(req *dto.UpdateDosenNamaReqDTO) error {
	dtDosen := assembler.ToUpdateDosenNama(req)

	err := s.repo.UpdateDosenNama(dtDosen)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) SaveDosenAlamatByID(req *dto.AlamatDosenByIDReqDTO) error {
	dtAlamatDOsen := assembler.ToSaveDosenAlamatByID(req)

	err := s.repo.SaveDosenAlamatByID(dtAlamatDOsen)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) ShowAllDosenAlamat(req *dto.DosenParamReqDTO) ([]*dto.DosenAlamatResDTO, error) {
	// dtDosen := assembler.ToShowAllDosenAlamat(req)

	var where string
	if req.IdDosen > 0 {
		where = fmt.Sprintf("a.id = %d", req.IdDosen)
	}
	if req.Nama != "" {
		where = fmt.Sprintf("a.nama = %s", req.Nama)
	}
	if req.Nidn != "" {
		where = fmt.Sprintf("a.nidn = %s", req.Nidn)
	}

	if req.IdDosen == 0 && req.Nama == "" && req.Nidn == "" {
		where = "1=1"
	}

	dataDosenAlamat, err := s.repo.ShowAllDosenAlamat(where)
	if err != nil {
		return nil, err
	}

	getDosensMap := make(map[int]*dto.DosenAlamatResDTO)
	for _, val := range dataDosenAlamat {
		if _, ok := getDosensMap[int(val.ID)]; !ok {
			getDosensMap[int(val.ID)] = &dto.DosenAlamatResDTO{
				ID:   val.ID,
				Nama: val.Name,
				Nidn: val.Nidn,
			}
			getDosensMap[int(val.ID)].Alamats = append(getDosensMap[int(val.ID)].Alamats, &dto.AlamatDosenResDTO{
				Jalan:   val.Jalan,
				NoRumah: val.NoRumah,
			})
		} else {
			getDosensMap[int(val.ID)].Alamats = append(getDosensMap[int(val.ID)].Alamats, &dto.AlamatDosenResDTO{
				Jalan:   val.Jalan,
				NoRumah: val.NoRumah,
			})
		}
	}

	var Data []*dto.DosenAlamatResDTO
	for _, datas := range getDosensMap {
		Data = append(Data, datas)
	}

	return Data, nil
}

func (s *service) GetMahasiswaAlamat(req *dto.GetMahasiswaAlamatReqDTO) ([]*dto.GetMahasiswaAlamatRespDTO, error) {
	resp := []*dto.GetMahasiswaAlamatRespDTO{}
	query := "%s"
	where := `'%s'`
	if req.Nama != "" {
		where = fmt.Sprintf(where, req.Nama)
		query = fmt.Sprintf(query, `a.nama = `+where)
	}

	getMahasiswaMap := make(map[int64]*dto.GetMahasiswaAlamatRespDTO)

	data, err := s.repo.GetMahasiswaAlamat(query)

	if err != nil {
		return nil, err
	}

	for _, val := range data {
		if _, ok := getMahasiswaMap[val.ID]; !ok {
			getMahasiswaMap[val.ID] = &dto.GetMahasiswaAlamatRespDTO{
				ID:   val.ID,
				Nama: val.Name,
				Nim:  val.Nim,
			}
			getMahasiswaMap[val.ID].Alamats = append(getMahasiswaMap[val.ID].Alamats, &dto.AlamatRespDTO{
				Jalan:   val.Jalan,
				NoRumah: val.NoRumah,
			})
		} else {
			getMahasiswaMap[val.ID].Alamats = append(getMahasiswaMap[val.ID].Alamats, &dto.AlamatRespDTO{
				Jalan:   val.Jalan,
				NoRumah: val.NoRumah,
			})
		}
	}

	for _, val := range getMahasiswaMap {
		resp = append(resp, val)
	}

	return resp, nil

}

func (s *service) GetIntegDadJoke(req *dto.GetDadJokesInternalReqDTO) (*dto.GetDadJokesRandomRespDTO, error) {
	var resp *dto.GetDadJokesRandomRespDTO

	resp, err := s.IntegServ.GetRandomDadJokes(req)

	if err != nil {
		return nil, err
	}

	return resp, nil

}
