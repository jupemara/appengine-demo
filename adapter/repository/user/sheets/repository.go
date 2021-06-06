package sheets

import (
	"context"

	"github.com/jupemara/appengine-demo/domain/model/user"
	sheets "github.com/jupemara/go-spreadsheet-sql"
	"go.opentelemetry.io/otel"
)

type userRepository struct {
	client *sheets.Client
}

func NewUserRepository(client *sheets.Client) *userRepository {
	return &userRepository{client}
}

func (r userRepository) VisibleName() string {
	return "GOOGLE_SHEETS"
}

func (r *userRepository) List(ctx context.Context) ([]*user.User, error) {
	tr := otel.GetTracerProvider().Tracer("appengine-demo/list")
	_, span := tr.Start(ctx, r.VisibleName())
	defer span.End()
	rows, err := r.client.Query(
		context.TODO(),
		`SELECT *`,
	)
	if err != nil {
		return nil, err
	}
	vs := []schema{}
	if err := rows.DataTo(&vs); err != nil {
		return nil, err
	}
	results := make([]*user.User, 0, len(vs))
	for _, v := range vs {
		results = append(results, user.NewUser(
			v.Id,
			v.FamilyName,
			v.GivenName,
			v.Email,
		))
	}
	return results, nil
}
