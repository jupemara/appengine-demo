package sheets

import (
	"context"

	"github.com/jupemara/appengine-demo/domain/model/user"
	sheets "github.com/jupemara/go-spreadsheet-sql"
)

type userRepository struct {
	client *sheets.Client
}

func NewUserRepository(client *sheets.Client) *userRepository {
	return &userRepository{client}
}

func (r *userRepository) VisibleName() string {
	return "GOOGLE_SHEETS"
}

func (r *userRepository) List() ([]*user.User, error) {
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
