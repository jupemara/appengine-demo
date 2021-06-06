package user

import (
	"encoding/json"
	"net/http"

	usecase "github.com/jupemara/appengine-demo/usecase/user"
)

type httpUserListHandler struct {
	usecase *usecase.ListUserUsecase
}

type response struct {
	Samples []sample `json:"samples"`
}

type sample struct {
	DataSourceName string `json:"data_source_name"`
	Users          []user `json:"users"`
}

type user struct {
	Id         string `json:"id"`
	FamilyName string `json:"family_name"`
	GivenName  string `json:"given_name"`
	Email      string `json:"email"`
}

func NewHttpUserListHandler(usecase *usecase.ListUserUsecase) *httpUserListHandler {
	return &httpUserListHandler{usecase}
}

func (h *httpUserListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vs, err := h.usecase.Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	samples := make([]sample, 0, len(vs))
	for _, v := range vs {
		us := make([]user, 0, len(v.Users))
		for _, u := range v.Users {
			us = append(us, user{
				Id:         u.Id(),
				FamilyName: u.FamilyName(),
				GivenName:  u.GivenName(),
				Email:      u.Email(),
			})
		}
		samples = append(samples, sample{
			DataSourceName: v.Repository,
			Users:          us,
		})
	}
	w.Header().Set("Content-Type", "application/json")
	body, err := json.MarshalIndent(response{
		Samples: samples,
	}, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(body)
}