package handles

import (
	"encoding/json"
	"net/http"
	"time"

	"strconv"

	"github.com/labstack/echo/v4"
	"gitlab.com/charlestenorios/next_ride_backend/app/database"
	"gitlab.com/charlestenorios/next_ride_backend/app/repository"
	"gitlab.com/charlestenorios/next_ride_backend/app/service"
)

func CreateCysclis(c echo.Context) error {
	db, err := database.Conn()
	repositorySql := repository.NewCyclistRepositoryPsql(db)
	serv := service.NewCycistService(repositorySql)

	idGrupo := c.FormValue("idGrupo")
	name := c.FormValue("name")
	cpf := c.FormValue("cpf")
	birth := c.FormValue("birth")
	email := c.FormValue("email")
	bloodType := c.FormValue("bloodType")
	healthPlan := c.FormValue("healthPlan")
	contactEmergency := c.FormValue("contactEmergency")
	gotToKnow := c.FormValue("gotToKnow")
	img := c.FormValue("img")
	participantType := c.FormValue("participantType")
	birthDay, err := time.Parse("2006-01-02", birth)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	if name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	cycli, err := serv.Create(idGrupo, name, cpf, birthDay, email, bloodType, healthPlan,
		contactEmergency, gotToKnow, img, participantType)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	if err != nil {
		return c.String(http.StatusInternalServerError, string(err.Error()))
	}
	return c.JSON(http.StatusCreated, cycli)
}

func GetByIdCysclist(c echo.Context) error {
	db, err := database.Conn()

	if err != nil {
		return c.String(http.StatusInternalServerError, string(err.Error()))

	}
	repositorySql := repository.NewCyclistRepositoryPsql(db)
	serv := service.NewCycistService(repositorySql)

	id := c.FormValue("id")
	cyclist, err := serv.Repository.Get(id)

	if err != nil {
		return c.String(http.StatusInternalServerError, string(err.Error()))

	}

	cy, err := json.Marshal(cyclist)

	if err != nil {
		return c.String(http.StatusInternalServerError, string(err.Error()))

	}
	return c.JSON(http.StatusCreated, cy)
}

func GetAllCyclist(c echo.Context) error {
	db, err := database.Conn()
	repositorySql := repository.NewCyclistRepositoryPsql(db)
	serv := service.NewCycistService(repositorySql)

	cycist, err := serv.Repository.FindAll()

	if err != nil {
		return c.String(http.StatusInternalServerError, string(err.Error()))

	}

	cy, err := json.Marshal(cycist)

	if err != nil {
		return c.String(http.StatusInternalServerError, string(err.Error()))
	}
	return c.JSON(http.StatusOK, cy)
}

func UpdateCyclist(c echo.Context) error {
	db, err := database.Conn()
	repositorySql := repository.NewCyclistRepositoryPsql(db)
	serv := service.NewCycistService(repositorySql)

	id := c.FormValue("idCyclist")
	idGrupo := c.FormValue("idGrupo")
	name := c.FormValue("name")
	cpf := c.FormValue("cpf")
	birth := c.FormValue("birth")
	email := c.FormValue("email")
	bloodType := c.FormValue("bloodType")
	healthPlan := c.FormValue("healthPlan")
	contactEmergency := c.FormValue("contactEmergency")
	gotToKnow := c.FormValue("gotToKnow")
	img := c.FormValue("img")
	participantType := c.FormValue("participantType")
	pedaling := c.FormValue("pedaling ")
	tours := c.FormValue("tours")
	travels := c.FormValue("travels")
	active := c.FormValue("active")
	bActivate, err := strconv.ParseBool(active)

	if err != nil {
		return c.String(http.StatusInternalServerError, string(err.Error()))

	}

	birthDay, err := time.Parse("2006-01-02", birth)

	if err != nil {
		return c.String(http.StatusInternalServerError, string(err.Error()))
	}

	if name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	intPedaling, err := strconv.Atoi(pedaling)

	if err != nil {
		return c.String(http.StatusInternalServerError, string(err.Error()))
	}

	intTours, err := strconv.Atoi(tours)

	if err != nil {
		return c.String(http.StatusInternalServerError, string(err.Error()))
	}

	intTravel, err := strconv.Atoi(travels)

	if err != nil {

		return c.String(http.StatusInternalServerError, string(err.Error()))

	}

	cycli, err := serv.UpdateCyclist(id, idGrupo, name, cpf, birthDay, email, bloodType, healthPlan, contactEmergency,
		gotToKnow, bActivate, img, participantType, intPedaling, intTours, intTravel)
	if err != nil {

		return c.String(http.StatusInternalServerError, string(err.Error()))

	}

	b, err := json.Marshal(cycli)

	if err != nil {
		return c.String(http.StatusInternalServerError, string(err.Error()))

	}

	return c.JSON(http.StatusCreated, b)
}

func DeleteCyclist(c echo.Context) error {
	db, err := database.Conn()
	repositorySql := repository.NewCyclistRepositoryPsql(db)
	serv := service.NewCycistService(repositorySql)

	id := c.FormValue("id")
	grup, _ := serv.Delete(id)

	g, err := json.Marshal(grup)

	if err != nil {
		return c.String(http.StatusInternalServerError, string(err.Error()))
	}
	return c.JSON(http.StatusAccepted, g)
}

func GetByNameCysclist(c echo.Context) error {
	db, err := database.Conn()

	if err != nil {
		return c.String(http.StatusInternalServerError, string(err.Error()))

	}
	repositorySql := repository.NewCyclistRepositoryPsql(db)
	serv := service.NewCycistService(repositorySql)
	name := c.FormValue("name")
	cyclist, err := serv.Repository.GetByName(name)

	if err != nil {
		return c.String(http.StatusInternalServerError, string(err.Error()))

	}

	cy, err := json.Marshal(cyclist)

	if err != nil {
		return c.String(http.StatusInternalServerError, string(err.Error()))

	}
	return c.JSON(http.StatusFound, cy)
}

func GetByCpfCysclist(c echo.Context) error {
	db, err := database.Conn()

	if err != nil {
		return c.String(http.StatusInternalServerError, string(err.Error()))

	}
	repositorySql := repository.NewCyclistRepositoryPsql(db)
	serv := service.NewCycistService(repositorySql)
	cpf := c.FormValue("cpf")
	cyclist, err := serv.Repository.GetByCpf(cpf)

	if err != nil {
		return c.String(http.StatusInternalServerError, string(err.Error()))

	}

	cy, err := json.Marshal(cyclist)

	if err != nil {
		return c.String(http.StatusInternalServerError, string(err.Error()))

	}
	return c.JSON(http.StatusFound, cy)
}

func UpdatePedaling(c echo.Context) error {
	db, err := database.Conn()
	repositorySql := repository.NewCyclistRepositoryPsql(db)
	serv := service.NewCycistService(repositorySql)
	id := c.FormValue("idCyclist")
	pedaling := c.FormValue("pedaling")
	intPedaling, err := strconv.Atoi(pedaling)

	if err != nil {
		return c.String(http.StatusInternalServerError, string(err.Error()))
	}

	cycli, err := serv.UpdatePedaling(id, intPedaling)

	if err != nil {
		return c.String(http.StatusInternalServerError, string(err.Error()))

	}

	b, err := json.Marshal(cycli)

	if err != nil {
		return c.String(http.StatusInternalServerError, string(err.Error()))

	}

	return c.JSON(http.StatusCreated, b)
}

func UpdateTravels(c echo.Context) error {
	db, err := database.Conn()
	repositorySql := repository.NewCyclistRepositoryPsql(db)
	serv := service.NewCycistService(repositorySql)
	id := c.FormValue("idCyclist")
	tour := c.FormValue("tour")
	intTour, err := strconv.Atoi(tour)

	if err != nil {
		return c.String(http.StatusInternalServerError, string(err.Error()))
	}

	cycli, err := serv.UpdateTours(id, intTour)

	if err != nil {
		return c.String(http.StatusInternalServerError, string(err.Error()))

	}

	b, err := json.Marshal(cycli)

	if err != nil {
		return c.String(http.StatusInternalServerError, string(err.Error()))

	}

	return c.JSON(http.StatusCreated, b)
}

func UpdateTours(c echo.Context) error {
	db, err := database.Conn()
	repositorySql := repository.NewCyclistRepositoryPsql(db)
	serv := service.NewCycistService(repositorySql)
	id := c.FormValue("idCyclist")
	travel := c.FormValue("tour")
	intTravel, err := strconv.Atoi(travel)
	if err != nil {
		return c.String(http.StatusInternalServerError, string(err.Error()))
	}

	cycli, err := serv.UpdateTravels(id, intTravel)

	if err != nil {
		return c.String(http.StatusInternalServerError, string(err.Error()))

	}

	b, err := json.Marshal(cycli)

	if err != nil {
		return c.String(http.StatusInternalServerError, string(err.Error()))

	}

	return c.JSON(http.StatusCreated, b)
}
