package controllers

import (
	"gitlab.com/charlestenorios/next_ride_backend/app/domain"
	"gitlab.com/charlestenorios/next_ride_backend/app/dtos"
	"gitlab.com/charlestenorios/next_ride_backend/app/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func FileUpload(c echo.Context) error {
	//upload
	formHeader, err := c.FormFile("file")
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			dtos.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &echo.Map{"data": "Select a file to upload"},
			})
	}

	//get file from header
	formFile, err := formHeader.Open()
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			dtos.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &echo.Map{"data": err.Error()},
			})
	}

	uploadUrl, err := service.NewMediaUpload().FileUpload(domain.File{File: formFile})
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			dtos.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &echo.Map{"data": err.Error()},
			})
	}

	return c.JSON(
		http.StatusOK,
		dtos.MediaDto{
			StatusCode: http.StatusOK,
			Message:    "success",
			Data:       &echo.Map{"data": uploadUrl},
		})
}

func RemoteUpload(c echo.Context) error {
	var url domain.Url

	//validate the request body
	if err := c.Bind(&url); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			dtos.MediaDto{
				StatusCode: http.StatusBadRequest,
				Message:    "error",
				Data:       &echo.Map{"data": err.Error()},
			})
	}

	uploadUrl, err := service.NewMediaUpload().RemoteUpload(url)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			dtos.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &echo.Map{"data": "Error uploading file"},
			})
	}

	return c.JSON(
		http.StatusOK,
		dtos.MediaDto{
			StatusCode: http.StatusOK,
			Message:    "success",
			Data:       &echo.Map{"data": uploadUrl},
		})
}
