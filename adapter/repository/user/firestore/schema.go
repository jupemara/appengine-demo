package firestore

type schema struct {
	Id         string `firestore:"id"`
	FamilyName string `firestore:"family_name"`
	GivenName  string `firestore:"given_name"`
	Email      string `firestore:"email"`
}
