package http

import (
	"encoding/json"
	"net/http"

	"github.com/eriktate/wrkhub"
	"github.com/sirupsen/logrus"
)

func GetCurrentUser(service wrkhub.UserService, log *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := GetSession(r.Context())
		if err != nil {
			log.WithError(err).Error("could not extract session data")
			forbidden(w, "invalid session")
			return
		}

		user, err := service.FetchUser(r.Context(), session.UserID)
		if err != nil {
			log.WithError(err).Error("could not fetch user")
			serverError(w, "could not fetch user")
			return
		}

		data, err := json.Marshal(user)
		if err != nil {
			log.WithError(err).Error("could not marshal user")
			serverError(w, "could not fetch user")
			return
		}

		ok(w, data)
	}
}
