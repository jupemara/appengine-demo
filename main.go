package main

import (
	"log"
	"net/http"
	"os"

	handler "github.com/jupemara/appengine-demo/adapter/controller/http/user"
	"github.com/jupemara/appengine-demo/adapter/repository/user/csv"
	domain "github.com/jupemara/appengine-demo/domain/model/user"
	usecase "github.com/jupemara/appengine-demo/usecase/user"
)

func main() {
	mux := http.NewServeMux()
	listHandler := handler.NewHttpUserListHandler(
		usecase.NewListUserUsecase(
			[]domain.UserRepository{
				csv.NewUserRepository("./data/users.csv"),
			},
		),
	)
	mux.Handle("/list", listHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}
