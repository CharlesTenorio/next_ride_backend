package handles

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.com/charlestenorios/next_ride_backend/app/database"
	"gitlab.com/charlestenorios/next_ride_backend/app/repository"
	"gitlab.com/charlestenorios/next_ride_backend/app/service"
)

var db, err = database.Conn()
var repositorySql1 = repository.NewGroupRepositoryPsql(db)
var serv = service.NewGroupService(repositorySql1)

func CreateGroup(c echo.Context) error {
	name := c.FormValue("name")
	if name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	grup, _ := serv.Create(name)
	g, err := json.Marshal(grup)

	if err != nil {
		return c.String(http.StatusInternalServerError, string(err.Error()))
	}
	return c.JSON(http.StatusCreated, g)
}

func GetByIdGroup(c echo.Context) error {
	id := c.FormValue("id")
	grup, _ := serv.Repository.Get(id)

	g, err := json.Marshal(grup)

	if err != nil {
		return c.String(http.StatusInternalServerError, string(err.Error()))

	}
	return c.JSON(http.StatusCreated, g)
}

func GetAllGroup(c echo.Context) error {
	grup, _ := serv.Repository.FindAll()
	g, err := json.Marshal(grup)

	if err != nil {
		return c.String(http.StatusInternalServerError, string(err.Error()))
	}
	return c.JSON(http.StatusOK, g)
}

func UpdateGroup(c echo.Context) error {
	id := c.FormValue("id")
	name := c.FormValue("name")

	grup, _ := serv.Update(id, name)

	g, err := json.Marshal(grup)

	if err != nil {
		return c.String(http.StatusInternalServerError, string(err.Error()))

	}
	return c.JSON(http.StatusCreated, g)
}

func DeleteGroup(c echo.Context) error {
	id := c.FormValue("id")
	grup, _ := serv.Delete(id)

	g, err := json.Marshal(grup)

	if err != nil {
		return c.String(http.StatusInternalServerError, string(err.Error()))
	}
	return c.JSON(http.StatusAccepted, g)
}
