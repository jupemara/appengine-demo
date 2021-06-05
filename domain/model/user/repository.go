package user

type UserRepository interface {
	List() []*user
}
