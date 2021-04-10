package http

import (
	"github.com/asdamwongmantap/api-gomongo-crud/crud"
	"github.com/asdamwongmantap/api-gomongo-crud/crud/delivery/controller"
	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, crudUseCase crud.CrudUseCaseI) {

	crudCtrl := controller.NewCrudController(e, crudUseCase)

	r := e.Group("/api/v1/go-mongo")
	r.GET("/list", crudCtrl.GetData)

}
