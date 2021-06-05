package user

type user struct {
	id         string
	familyName string
	givenName  string
	email      string
}

func NewUser(
	id,
	familyName,
	givenName,
	email string,
) *user {
	return &user{id, familyName, givenName, email}
}

func (u *user) Id() string {
	return u.id
}

func (u *user) FamilyName() string {
	return u.familyName
}

func (u *user) GivenName() string {
	return u.givenName
}

func (u *user) Email() string {
	return u.email
}
