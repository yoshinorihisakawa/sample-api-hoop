package router

import (
	"github.com/labstack/echo"
	"github.com/yoshinorihisakawa/sample-api-hoop/infrastructure/api/handler"
)

func NewRouter(e *echo.Echo, handler handler.AppHandler) {
	e.POST("/users", handler.CreateUser)
	e.GET("/users", handler.GetUsers)
}
