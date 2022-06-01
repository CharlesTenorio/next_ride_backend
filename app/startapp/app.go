package startapp

import (
	"fmt"
	"gitlab.com/charlestenorios/next_ride_backend/app/controllers"
	"gitlab.com/charlestenorios/next_ride_backend/app/service"
	"log"
	"net/http"

	"encoding/json"

	"github.com/labstack/echo/v4"
	"gitlab.com/charlestenorios/next_ride_backend/app/database"
	"gitlab.com/charlestenorios/next_ride_backend/app/repository"
)

func Start() {
	//db := repository.GroupMemoryDb{Groups: []domain.Group{}}
	//repositoryMemory := repository.NewGroupRepositoryMemory(db)
	db, err := database.Conn()

	e := echo.New()
	e.POST("/criar_grupo", func(c echo.Context) error {

		if err != nil {
			fmt.Print("deu merda na conexao", string(err.Error()))
			log.Fatal("deu mreda na conexao do banco")
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
			fmt.Print(err.Error())

		}
		return c.String(http.StatusOK, string(b))
	})

	e.GET("/listar", func(c echo.Context) error {

		if err != nil {
			fmt.Print("deu merda na conexao", string(err.Error()))
			log.Fatal("deu mreda na conexao do banco")
		}

		repositorySql := repository.NewGroupRepositoryPsql(db)
		serv := service.NewGroupService(repositorySql)
		id := c.FormValue("id")
		grup, _ := serv.Repository.Get(id)

		b, err := json.Marshal(grup)

		if err != nil {
			fmt.Print(err.Error())

		}
		return c.String(http.StatusOK, string(b))
	})

	e.GET("/all", func(c echo.Context) error {

		if err != nil {
			fmt.Print("deu merda na conexao", string(err.Error()))
			log.Fatal("deu mreda na conexao do banco")
		}

		repositorySql := repository.NewGroupRepositoryPsql(db)
		serv := service.NewGroupService(repositorySql)
		grup, _ := serv.Repository.FindAll()

		b, err := json.Marshal(grup)

		if err != nil {
			fmt.Print(err.Error())

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
			fmt.Print(err.Error())

		}
		return c.String(http.StatusOK, string(b))
	})

	e.DELETE("/delete", func(c echo.Context) error {

		if err != nil {
			fmt.Print("deu merda na conexao", string(err.Error()))
			log.Fatal("deu mreda na conexao do banco")
		}

		repositorySql := repository.NewGroupRepositoryPsql(db)
		serv := service.NewGroupService(repositorySql)
		id := c.FormValue("id")
		grup, _ := serv.Delete(id)

		b, err := json.Marshal(grup)

		if err != nil {
			fmt.Print(err.Error())

		}
		return c.String(http.StatusOK, string(b))
	})

	e.POST("/file", controllers.FileUpload)
	e.POST("/remote", controllers.RemoteUpload)

	e.Logger.Fatal(e.Start(":1323"))

}
