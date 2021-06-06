package sheets

type schema struct {
	Id         string `json:"id"`
	FamilyName string `json:"family_name"`
	GivenName  string `json:"given_name"`
	Email      string `json:"email"`
}
