package csv

import (
	"context"
	"encoding/csv"
	"io"
	"os"

	"github.com/jupemara/appengine-demo/domain/model/user"
	"go.opentelemetry.io/otel"
)

type userRepository struct {
	db string
}

func NewUserRepository(db string) *userRepository {
	return &userRepository{db}
}

func (r userRepository) VisibleName() string {
	return "CSV"
}

func (r *userRepository) List(ctx context.Context) ([]*user.User, error) {
	tr := otel.GetTracerProvider().Tracer("appengine-demo/list")
	_, span := tr.Start(ctx, r.VisibleName())
	defer span.End()
	file, err := os.OpenFile(r.db, os.O_RDONLY, 0600)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(file)
	us := []*user.User{}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		us = append(us, user.NewUser(
			record[0],
			record[1],
			record[2],
			record[3],
		))
	}
	return us, nil
}
