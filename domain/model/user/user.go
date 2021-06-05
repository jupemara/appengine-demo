package user

type User struct {
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
) *User {
	return &User{id, familyName, givenName, email}
}

func (u *User) Id() string {
	return u.id
}

func (u *User) FamilyName() string {
	return u.familyName
}

func (u *User) GivenName() string {
	return u.givenName
}

func (u *User) Email() string {
	return u.email
}
