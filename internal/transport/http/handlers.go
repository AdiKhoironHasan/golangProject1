package http

import (
	"fmt"
	"net/http"
	"os"
	"reflect"

	"github.com/AdiKhoironHasan/golangProject1/internal/services"
	mhsConst "github.com/AdiKhoironHasan/golangProject1/pkg/common/const"
	"github.com/AdiKhoironHasan/golangProject1/pkg/dto"
	mhsErrors "github.com/AdiKhoironHasan/golangProject1/pkg/errors"
	"github.com/apex/log"

	"github.com/labstack/echo"
)

type HttpHandler struct {
	service services.Services
}

func NewHttpHandler(e *echo.Echo, srv services.Services) {
	handler := &HttpHandler{
		srv,
	}
	e.GET("api/v1/latihan/ping", handler.Ping)
	e.POST("api/v1/latihan/mahasiswa-alamat", handler.SaveMahasiswaAlamat)
	e.PATCH("api/v1/latihan/mahasiswa", handler.UpdateMahasiswaNama)
	e.POST("api/v1/latihan/alamat", handler.SaveAlamatId)
	// e.GET("api/v1/latihan/mahasiswa-alamat", handler.ShowAllMahasiswaAlamat)
	e.GET("api/v1/latihan/mahasiswa-alamat", handler.GetMahasiswaAlamat)

	e.POST("api/v1/latihan/dosen-alamat", handler.SaveDosenAlamat)
	e.PATCH("api/v1/latihan/dosen", handler.UpdateDosenNama)
	e.POST("api/v1/latihan/alamat-dosen", handler.SaveDosenAlamatByID)
	e.GET("api/v1/latihan/dosen-alamat", handler.ShowAllDosenAlamat)

	e.GET("api/v1/latihan/dad-jokes", handler.GetRandomDadJokes)
}

func (h *HttpHandler) Ping(c echo.Context) error {

	version := os.Getenv("VERSION")
	if version == "" {
		version = "pong"
	}

	data := version

	return c.JSON(http.StatusOK, data)

}

func (h *HttpHandler) UpdateMahasiswaNama(c echo.Context) error {
	patchDTO := dto.UpadeMahasiswaNamaReqDTO{}
	fmt.Println("type : ", reflect.TypeOf(patchDTO).Kind())
	if err := c.Bind(&patchDTO); err != nil { //bind = req ke variabel
		log.Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	err := patchDTO.Validate()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = h.service.UpdateMahasiswaNama(&patchDTO)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	var resp = dto.ResponseDTO{
		Success: true,
		Message: mhsConst.UpdateSuccess,
		Data:    nil,
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *HttpHandler) SaveAlamatId(c echo.Context) error {
	postDTO := dto.AlamatIdReqDTO{}
	if err := c.Bind(&postDTO); err != nil { //bind = req ke variabel
		log.Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}
	err := postDTO.Validate()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = h.service.SaveAlamatId(&postDTO)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	var resp = dto.ResponseDTO{
		Success: true,
		Message: mhsConst.SaveSuccess,
		Data:    nil,
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *HttpHandler) SaveMahasiswaAlamat(c echo.Context) error {
	postDTO := dto.MahasiswaReqDTO{}         //{} untuk memberi nilai default yaitu nil
	if err := c.Bind(&postDTO); err != nil { //bind = req ke variabel
		log.Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	err := postDTO.Validate()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = h.service.SaveMahasiswaAlamat(&postDTO)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	var resp = dto.ResponseDTO{
		Success: true,
		Message: mhsConst.SaveSuccess,
		Data:    nil,
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *HttpHandler) ShowAllMahasiswaAlamat(c echo.Context) error {
	Data, _ := h.service.ShowAllMahasiswaAlamat()

	var resp = dto.ResponseDTO{
		Success: true,
		Message: mhsConst.GetDataSuccess,
		Data:    Data,
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *HttpHandler) GetMahasiswaAlamat(c echo.Context) error {
	postDTO := dto.GetMahasiswaAlamatReqDTO{}

	postDTO.Authorization = c.Request().Header.Get("Authorization")
	postDTO.DateTime = c.Request().Header.Get("datetime")
	postDTO.Signature = c.Request().Header.Get("signature")

	if err := c.Bind(&postDTO); err != nil {
		log.Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	err := postDTO.Validate()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	data, err := h.service.GetMahasiswaAlamat(&postDTO)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	fmt.Print("data handler", data)
	var resp = dto.ResponseDTO{
		Success: true,
		Message: mhsConst.GetDataSuccess,
		Data:    data,
	}

	return c.JSON(http.StatusOK, resp)
}

// DOSEN
func (h *HttpHandler) SaveDosenAlamat(c echo.Context) error {
	postDTO := dto.DosenReqDTO{}

	if err := c.Bind(&postDTO); err != nil {
		log.Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	err := postDTO.Validate()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = h.service.SaveDosenAlamat(&postDTO)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	var resp = dto.ResponseDTO{
		Success: true,
		Message: mhsConst.SaveSuccess,
		Data:    nil,
	}

	return c.JSON(http.StatusOK, resp)
}

// DOSEN

func (h *HttpHandler) UpdateDosenNama(c echo.Context) error {
	patchDTO := dto.UpdateDosenNamaReqDTO{}
	if err := c.Bind(&patchDTO); err != nil {
		log.Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	err := patchDTO.Validate()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = h.service.UpdateDosenNama(&patchDTO)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	var resp = dto.ResponseDTO{
		Success: true,
		Message: mhsConst.UpdateSuccess,
		Data:    nil,
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *HttpHandler) SaveDosenAlamatByID(c echo.Context) error {
	postDTO := dto.AlamatDosenByIDReqDTO{}
	if err := c.Bind(&postDTO); err != nil {
		log.Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	err := postDTO.Validate()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = h.service.SaveDosenAlamatByID(&postDTO)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	var resp = dto.ResponseDTO{
		Success: true,
		Message: mhsConst.SaveSuccess,
		Data:    nil,
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *HttpHandler) ShowAllDosenAlamat(c echo.Context) error {
	getDTO := dto.DosenParamReqDTO{}
	if err := c.Bind(&getDTO); err != nil {
		log.Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	err := getDTO.Validate()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	result, err := h.service.ShowAllDosenAlamat(&getDTO)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	var resp = dto.ResponseDTO{
		Success: true,
		Message: mhsConst.GetDataSuccess,
		Data:    result,
	}

	return c.JSON(http.StatusOK, resp)

}

func (h *HttpHandler) GetRandomDadJokes(c echo.Context) error {
	postDTO := dto.GetDadJokesInternalReqDTO{}

	postDTO.Authorization = c.Request().Header.Get("Authorization")
	postDTO.DateTime = c.Request().Header.Get("datetime")
	postDTO.Signature = c.Request().Header.Get("signature")

	if err := c.Bind(&postDTO); err != nil {
		log.Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	err := postDTO.Validate()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	// fmt.Println("dto: ", postDTO.ID)

	data, err := h.service.GetIntegDadJoke(&postDTO)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	var resp = dto.ResponseDTO{
		Success: true,
		Message: mhsConst.GetDataSuccess,
		Data:    data,
	}

	return c.JSON(http.StatusOK, resp)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case mhsErrors.ErrInternalServerError:
		return http.StatusInternalServerError
	case mhsErrors.ErrNotFound:
		return http.StatusNotFound
	case mhsErrors.ErrConflict:
		return http.StatusConflict
	case mhsErrors.ErrInvalidRequest:
		return http.StatusBadRequest
	case mhsErrors.ErrFailAuth:
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}
