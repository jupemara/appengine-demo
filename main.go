package main

import (
	"context"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	"github.com/jupemara/appengine-demo/adapter/config"
	handler "github.com/jupemara/appengine-demo/adapter/controller/http/user"
	"github.com/jupemara/appengine-demo/adapter/repository/user/csv"
	"github.com/jupemara/appengine-demo/adapter/repository/user/firestore"
	"github.com/jupemara/appengine-demo/adapter/repository/user/sheets"
	domain "github.com/jupemara/appengine-demo/domain/model/user"
	usecase "github.com/jupemara/appengine-demo/usecase/user"
	spreadsheets "github.com/jupemara/go-spreadsheet-sql"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
)

func main() {
	const e = `unexpected error occurred: `
	// load application config
	c := config.NewConfig()
	if err := c.Load(); err != nil {
		log.Fatalf(`%s%s`, e, err)
	}

	// cloud trace + open telemetry setup
	ctx := context.TODO()
	exp, err := texporter.NewExporter(texporter.WithProjectID(c.GcpProjectId()))
	if err != nil {
		log.Fatalf(`%s%s`, e, err)
	}
	defer exp.Shutdown(ctx)
	tp := trace.NewTracerProvider(trace.WithSyncer(exp))
	otel.SetTracerProvider(tp)

	// google sheets setup
	sheetsClient, err := spreadsheets.NewClient(
		context.TODO(),
		c.SheetsKey(),
		c.SheetsName(),
	)
	if err != nil {
		log.Fatalf(`%s%s`, e, err)
	}
	// firestore client setup
	firebaseApp, err := firebase.NewApp(
		context.TODO(),
		&firebase.Config{ProjectID: c.GcpProjectId()},
	)
	if err != nil {
		log.Fatalf(`%s%s`, e, err)
	}
	firestoreClient, err := firebaseApp.Firestore(context.TODO())
	if err != nil {
		log.Fatalf(`%s%s`, e, err)
	}

	// DI: list users handler
	listHandler := handler.NewHttpUserListHandler(
		usecase.NewListUserUsecase(
			[]domain.UserRepository{
				csv.NewUserRepository("./data/users.csv"),
				sheets.NewUserRepository(sheetsClient),
				firestore.NewUserRepository(firestoreClient),
			},
		),
	)
	mux := http.NewServeMux()
	mux.Handle("/", listHandler)
	if err := http.ListenAndServe(":"+c.Port(), mux); err != nil {
		log.Fatal(err)
	}
}
