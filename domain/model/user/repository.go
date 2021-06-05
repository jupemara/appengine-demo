package user

type UserRepository interface {
	List() ([]*User, error)
	VisibleName() string
}
