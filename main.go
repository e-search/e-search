package main

import (
	"database/sql"
	"log"

	"github.com/e-search/e-search/account/controller"
	"github.com/e-search/e-search/account/repository"
	"github.com/e-search/e-search/account/usecase"
	"github.com/e-search/e-search/toml"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

func main() {
	url, err := toml.LoadDBToml()
	if err != nil {
		log.Println(err)
	}
	dbConn, err := sql.Open("postgres", url)
	if err != nil {
		log.Println(err)
	}
	defer dbConn.Close()
	e := echo.New()

	accountRepo := repository.NewAccountRepository(dbConn)
	accountUsecase := usecase.NewAccountUsecase(accountRepo)

	controller.NewAccountController(e, accountUsecase)

	e.Logger.Fatal(e.Start(":8080"))
}
