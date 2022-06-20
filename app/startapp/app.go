package startapp

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"gitlab.com/charlestenorios/next_ride_backend/app/controllers"
	"gitlab.com/charlestenorios/next_ride_backend/app/service"

	"encoding/json"

	"github.com/labstack/echo/v4"
	"gitlab.com/charlestenorios/next_ride_backend/app/database"
	"gitlab.com/charlestenorios/next_ride_backend/app/repository"
)

func Start() {

	db, err := database.Conn()

	e := echo.New()
	e.POST("/create_group", func(c echo.Context) error {

		if err != nil {
			fmt.Print("deu merda na conexao", string(err.Error()))
			return c.String(http.StatusInternalServerError, string(err.Error()))
		}

		repositorySql := repository.NewGroupRepositoryPsql(db)
		serv := service.NewGroupService(repositorySql)
		name := c.FormValue("name")

		if name == "" {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		grup, _ := serv.Create(name)

		b, err := json.Marshal(grup)

		if err != nil {
			return c.String(http.StatusInternalServerError, string(err.Error()))

		}
		return c.String(http.StatusOK, string(b))
	})

	e.GET("/list_by_id_group", func(c echo.Context) error {

		if err != nil {
			return c.String(http.StatusInternalServerError, string(err.Error()))
		}

		repositorySql := repository.NewGroupRepositoryPsql(db)
		serv := service.NewGroupService(repositorySql)
		id := c.FormValue("id")
		grup, _ := serv.Repository.Get(id)

		b, err := json.Marshal(grup)

		if err != nil {
			return c.String(http.StatusInternalServerError, string(err.Error()))

		}
		return c.String(http.StatusOK, string(b))
	})

	e.GET("/all", func(c echo.Context) error {

		if err != nil {
			return c.String(http.StatusInternalServerError, string(err.Error()))
		}

		repositorySql := repository.NewGroupRepositoryPsql(db)
		serv := service.NewGroupService(repositorySql)
		grup, _ := serv.Repository.FindAll()

		b, err := json.Marshal(grup)

		if err != nil {
			return c.String(http.StatusInternalServerError, string(err.Error()))

		}
		return c.String(http.StatusOK, string(b))
	})

	e.POST("/update", func(c echo.Context) error {

		if err != nil {
			log.Fatal("deu ruim na conexao do banco")
		}

		repositorySql := repository.NewGroupRepositoryPsql(db)
		serv := service.NewGroupService(repositorySql)
		id := c.FormValue("id")
		name := c.FormValue("name")

		grup, _ := serv.Update(id, name)

		b, err := json.Marshal(grup)

		if err != nil {
			return c.String(http.StatusInternalServerError, string(err.Error()))

		}
		return c.String(http.StatusOK, string(b))
	})

	e.DELETE("/delete", func(c echo.Context) error {

		if err != nil {
			return c.String(http.StatusInternalServerError, string(err.Error()))
		}

		repositorySql := repository.NewGroupRepositoryPsql(db)
		serv := service.NewGroupService(repositorySql)
		id := c.FormValue("id")
		grup, _ := serv.Delete(id)

		b, err := json.Marshal(grup)

		if err != nil {
			return c.String(http.StatusInternalServerError, string(err.Error()))

		}
		return c.String(http.StatusOK, string(b))
	})

	// endpoints ciclysts
	e.POST("/create_cyclist", func(c echo.Context) error {

		if err != nil {
			fmt.Print("deu merda na conexao", string(err.Error()))
			log.Fatal("deu mreda na conexao do banco")
		}

		cyclistRepositorySql := repository.NewCyclistRepositoryPsql(db)
		serv := service.NewCycistService(cyclistRepositorySql)
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
			fmt.Println(err)
		}

		if name == "" {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		cycli, _ := serv.Create(idGrupo, name, cpf, birthDay, email, bloodType, healthPlan,
			contactEmergency, gotToKnow, img, participantType)

		b, err := json.Marshal(cycli)

		if err != nil {
			fmt.Print(err.Error())

		}
		return c.String(http.StatusOK, string(b))
	})

	e.POST("/update_cyclist", func(c echo.Context) error {

		if err != nil {
			return c.String(http.StatusInternalServerError, string(err.Error()))
		}

		cyclistRepositorySql := repository.NewCyclistRepositoryPsql(db)
		serv := service.NewCycistService(cyclistRepositorySql)
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
		bActivate, errc := strconv.ParseBool(active)
		if errc != nil {
			fmt.Println(err)
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
			fmt.Println(err)
		}

		cycli, _ := serv.UpdateCyclist(id, idGrupo, name, cpf, birthDay, email, bloodType, healthPlan, contactEmergency,
			gotToKnow, bActivate, img, participantType, intPedaling, intTours, intTravel)

		b, err := json.Marshal(cycli)

		if err != nil {
			return c.String(http.StatusInternalServerError, string(err.Error()))

		}
		return c.String(http.StatusOK, string(b))
	})

	e.GET("/all_cyclist", func(c echo.Context) error {

		if err != nil {
			fmt.Print("deu merda na conexao", string(err.Error()))
			log.Fatal("deu mreda na conexao do banco")
		}

		cyclistRepositorySql := repository.NewCyclistRepositoryPsql(db)
		serv := service.NewCycistService(cyclistRepositorySql)
		cyclists, _ := serv.Repository.FindAll()

		b, err := json.Marshal(cyclists)

		if err != nil {
			fmt.Print(err.Error())

		}
		return c.String(http.StatusOK, string(b))
	})

	e.GET("/find_by_id_cytlist", func(c echo.Context) error {

		if err != nil {
			fmt.Print("deu merda na conexao", string(err.Error()))
			log.Fatal("deu mreda na conexao do banco")
		}

		cyclistRepositorySql := repository.NewCyclistRepositoryPsql(db)
		serv := service.NewCycistService(cyclistRepositorySql)
		id := c.FormValue("id")
		cyclist, _ := serv.Repository.Get(id)

		b, err := json.Marshal(cyclist)

		if err != nil {
			fmt.Print(err.Error())

		}
		return c.String(http.StatusOK, string(b))
	})

	e.GET("/find_by_name_cytlist", func(c echo.Context) error {

		if err != nil {
			fmt.Print("deu merda na conexao", string(err.Error()))
			log.Fatal("deu mreda na conexao do banco")
		}

		cyclistRepositorySql := repository.NewCyclistRepositoryPsql(db)
		serv := service.NewCycistService(cyclistRepositorySql)
		name := c.FormValue("name")
		cyclist, _ := serv.Repository.GetByName(name)

		b, err := json.Marshal(cyclist)

		if err != nil {
			fmt.Print(err.Error())

		}
		return c.String(http.StatusOK, string(b))
	})

	e.GET("/find_by_cpf_cytlist", func(c echo.Context) error {

		if err != nil {
			fmt.Print("deu merda na conexao", string(err.Error()))
			log.Fatal("deu mreda na conexao do banco")
		}

		cyclistRepositorySql := repository.NewCyclistRepositoryPsql(db)
		serv := service.NewCycistService(cyclistRepositorySql)
		cpf := c.FormValue("name")
		cyclist, _ := serv.Repository.GetByCpf(cpf)

		b, err := json.Marshal(cyclist)

		if err != nil {
			fmt.Print(err.Error())

		}
		return c.String(http.StatusOK, string(b))
	})

	e.DELETE("/delete_cyclist", func(c echo.Context) error {

		if err != nil {
			fmt.Print("deu merda na conexao", string(err.Error()))
			log.Fatal("deu mreda na conexao do banco")
		}

		cyclistRepositorySql := repository.NewCyclistRepositoryPsql(db)
		serv := service.NewCycistService(cyclistRepositorySql)
		id := c.FormValue("id")
		cyclist, _ := serv.Delete(id)

		b, err := json.Marshal(cyclist)

		if err != nil {
			fmt.Print(err.Error())

		}
		return c.String(http.StatusOK, string(b))
	})

	e.PUT("/update_pedaling", func(c echo.Context) error {

		if err != nil {
			return c.String(http.StatusInternalServerError, string(err.Error()))
		}

		cyclistRepositorySql := repository.NewCyclistRepositoryPsql(db)
		serv := service.NewCycistService(cyclistRepositorySql)
		id := c.FormValue("idCyclist")
		pedaling := c.FormValue("pedaling ")
		intPedaling, err := strconv.Atoi(pedaling)

		if err != nil {
			return c.String(http.StatusInternalServerError, string(err.Error()))
		}

		cycli, _ := serv.UpdatePedaling(id, intPedaling)

		b, err := json.Marshal(cycli)

		if err != nil {
			return c.String(http.StatusInternalServerError, string(err.Error()))

		}
		return c.String(http.StatusOK, string(b))
	})

	e.PUT("/update_tour", func(c echo.Context) error {

		if err != nil {
			return c.String(http.StatusInternalServerError, string(err.Error()))
		}

		cyclistRepositorySql := repository.NewCyclistRepositoryPsql(db)
		serv := service.NewCycistService(cyclistRepositorySql)
		id := c.FormValue("idCyclist")
		tour := c.FormValue("tour")
		intTour, err := strconv.Atoi(tour)

		if err != nil {
			return c.String(http.StatusInternalServerError, string(err.Error()))
		}

		cycli, _ := serv.UpdateTours(id, intTour)

		b, err := json.Marshal(cycli)

		if err != nil {
			return c.String(http.StatusInternalServerError, string(err.Error()))

		}
		return c.String(http.StatusOK, string(b))
	})

	e.PUT("/update_travel", func(c echo.Context) error {

		if err != nil {
			return c.String(http.StatusInternalServerError, string(err.Error()))
		}

		cyclistRepositorySql := repository.NewCyclistRepositoryPsql(db)
		serv := service.NewCycistService(cyclistRepositorySql)
		id := c.FormValue("idCyclist")
		travel := c.FormValue("tour")
		intTravel, err := strconv.Atoi(travel)

		if err != nil {
			return c.String(http.StatusInternalServerError, string(err.Error()))
		}

		cycli, _ := serv.UpdateTours(id, intTravel)

		b, err := json.Marshal(cycli)

		if err != nil {
			return c.String(http.StatusInternalServerError, string(err.Error()))

		}
		return c.String(http.StatusOK, string(b))
	})

	// end cicystes

	e.POST("/file", controllers.FileUpload)
	e.POST("/remote", controllers.RemoteUpload)

	e.Logger.Fatal(e.Start(":3000"))

}
