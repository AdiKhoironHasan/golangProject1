package dto

import (
	"errors"

	"github.com/AdiKhoironHasan/golangProject1/pkg/common/crypto"
	"github.com/AdiKhoironHasan/golangProject1/pkg/common/validator"
	"github.com/AdiKhoironHasan/golangProject1/pkg/env"
	util "github.com/AdiKhoironHasan/golangProject1/pkg/utils"
)

type GetDadJokesInternalReqDTO struct {
	Authorization string `json:"Authorization" valid:"required" validname:"authorization"`
	Signature     string `json:"signature" valid:"required" validname:"signature"`
	DateTime      string `json:"datetime" valid:"required" validname:"datetime"`
	ID            string `json:"ID,omitempty,string" query:"ID"`
}

func (dto *GetDadJokesInternalReqDTO) Validate() error {
	v := validator.NewValidate(dto)
	v.SetCustomValidation(true, func() error {
		return dto.customValidation()
	})
	return v.Validate()
}

func (dto *GetDadJokesInternalReqDTO) customValidation() error {

	signature := crypto.EncodeSHA256HMAC(util.GetBTBPrivKeySignature(), dto.Authorization, dto.DateTime)
	if signature != dto.Signature {
		if env.IsProduction() {
			return errors.New("invalid signature")
		}
		return errors.New("invalid signature" + " --> " + signature)
	}

	return nil
}

type GetDadJokesRandomRespDTO struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int64  `json:"status"`
}
