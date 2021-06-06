package main

import (
	"context"
	"log"
	"net/http"

	"github.com/jupemara/appengine-demo/adapter/config"
	handler "github.com/jupemara/appengine-demo/adapter/controller/http/user"
	"github.com/jupemara/appengine-demo/adapter/repository/user/csv"
	"github.com/jupemara/appengine-demo/adapter/repository/user/sheets"
	domain "github.com/jupemara/appengine-demo/domain/model/user"
	usecase "github.com/jupemara/appengine-demo/usecase/user"
	spreadsheets "github.com/jupemara/go-spreadsheet-sql"
)

func main() {
	const e = `unexpected error occurred: `
	c := config.NewConfig()
	if err := c.Load(); err != nil {
		log.Fatalf(`%s%s`, e, err)
	}
	mux := http.NewServeMux()
	sheetsClient, err := spreadsheets.NewClient(
		context.TODO(),
		c.SheetsKey(),
		c.SheetsName(),
	)
	if err != nil {
		log.Fatalf(`%s%s`, e, err)
	}
	listHandler := handler.NewHttpUserListHandler(
		usecase.NewListUserUsecase(
			[]domain.UserRepository{
				csv.NewUserRepository("./data/users.csv"),
				sheets.NewUserRepository(sheetsClient),
			},
		),
	)
	mux.Handle("/list", listHandler)
	if err := http.ListenAndServe(":"+c.Port(), mux); err != nil {
		log.Fatal(err)
	}
}
