package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/yoshinorihisakawa/sample-api-hoop/conf"
	"github.com/yoshinorihisakawa/sample-api-hoop/infrastructure/api/router"
	"github.com/yoshinorihisakawa/sample-api-hoop/infrastructure/api/validater"
	"github.com/yoshinorihisakawa/sample-api-hoop/infrastructure/datastore"
	"github.com/yoshinorihisakawa/sample-api-hoop/registry"
	"gopkg.in/go-playground/validator.v9"
)

func main() {
	// conf読み取り
	conf.ReadConf()

	// DB取得
	conn := datastore.NewMySqlDB()

	// interactor
	r := registry.NewInteractor(conn)

	// 依存解決
	h := r.NewAppHandler()

	// echo
	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// validate
	e.Validator = &validater.CustomValidator{Validator: validator.New()}

	// router
	router.NewRouter(e, h)

	// DB stop
	defer func() {
		if err := conn.Close(); err != nil {
			e.Logger.Fatal(fmt.Sprintf("Failed to close: %v", err))
		}
	}()

	// server start
	e.Logger.Fatal(e.Start(":" + conf.C.Server.Address))
}
