package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/jupemara/appengine-demo/domain/model/user"
	"go.opentelemetry.io/otel"
	"google.golang.org/api/iterator"
)

type userRepository struct {
	client *firestore.Client
}

func NewUserRepository(client *firestore.Client) *userRepository {
	return &userRepository{client}
}

func (r userRepository) VisibleName() string {
	return "CLOUD_FIRESTORE"
}

func (r *userRepository) List(ctx context.Context) ([]*user.User, error) {
	tr := otel.GetTracerProvider().Tracer("appengine-demo/list")
	_, span := tr.Start(ctx, r.VisibleName())
	defer span.End()
	fctx := context.TODO()
	ref := r.client.Collection("users").Documents(fctx)
	us := make([]*user.User, 0)
	for {
		doc, err := ref.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var u schema
		if err := doc.DataTo(&u); err != nil {
			return nil, err
		}
		us = append(us, user.NewUser(
			u.Id,
			u.FamilyName,
			u.GivenName,
			u.Email,
		))
	}
	return us, nil
}
