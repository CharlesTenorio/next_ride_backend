package service

import (
	"github.com/go-playground/validator/v10"
	"gitlab.com/charlestenorios/next_ride_backend/app/domain"
	"gitlab.com/charlestenorios/next_ride_backend/app/helper"
)

var (
	validate = validator.New()
)

type mediaUpload interface {
	FileUpload(file domain.File) (string, error)
	RemoteUpload(url domain.Url) (string, error)
}

type media struct{}

func NewMediaUpload() mediaUpload {
	return &media{}
}

func (*media) FileUpload(file domain.File) (string, error) {
	//validate
	err := validate.Struct(file)
	if err != nil {
		return "", err
	}

	//upload
	uploadUrl, err := helper.ImageUploadHelper(file.File)
	if err != nil {
		return "", err
	}
	return uploadUrl, nil
}

func (*media) RemoteUpload(url domain.Url) (string, error) {
	//validate
	err := validate.Struct(url)
	if err != nil {
		return "", err
	}

	//upload
	uploadUrl, errUrl := helper.ImageUploadHelper(url.Url)
	if errUrl != nil {
		return "", err
	}
	return uploadUrl, nil
}
