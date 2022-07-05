package startapp

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/charlestenorios/next_ride_backend/app/controllers"
	"gitlab.com/charlestenorios/next_ride_backend/app/handles"
)

func Start() {

	e := echo.New()
	e.POST("/create_group", handles.CreateGroup)
	e.GET("/get_id_group/:id", handles.GetByIdGroup)
	e.GET("/all", handles.GetAllGroup)
	e.PUT("/update/:id/name", handles.UpdateGroup)
	e.DELETE("/delete:id", handles.DeleteGroup)

	// endpoints ciclysts
	e.POST("/create_cyclist", handles.CreateCysclis)
	e.PUT("/update_cyclist", handles.UpdateCyclist)
	e.GET("/all_cyclist", handles.DeleteCyclist)
	e.GET("/find_by_id_cytlist/:id", handles.GetByIdCysclist)
	e.GET("/find_by_name_cytlist/:name", handles.GetByNameCysclist)
	e.GET("/find_by_cpf_cytlist/:cpf", handles.GetByCpfCysclist)
	e.DELETE("/delete_cyclist/:id", handles.DeleteGroup)
	e.PATCH("/update_pedaling/:id/pedaling", handles.UpdatePedaling)
	e.PATCH("/update_tour", handles.UpdateTours)
	e.PATCH("/update_travel/:id/tour", handles.UpdateTravels)

	// end cicystes

	e.POST("/file", controllers.FileUpload)
	e.POST("/remote", controllers.RemoteUpload)

	e.Logger.Fatal(e.Start(":3000"))

}
