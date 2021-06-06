package user

import "context"

type UserRepository interface {
	List(context.Context) ([]*User, error)
	VisibleName() string
}
