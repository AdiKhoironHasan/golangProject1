package http

import (
	"net/http"
	"os"

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
	postDTO := dto.UpadeMahasiswaNamaReqDTO{}
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

	err = h.service.UpdateMahasiswaNama(&postDTO)
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

func (h *HttpHandler) SaveMahasiswaAlamat(c echo.Context) error {
	postDTO := dto.MahasiswaReqDTO{}
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