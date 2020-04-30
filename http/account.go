package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/eriktate/uid"
	"github.com/eriktate/watdo"
	"github.com/go-chi/chi"
)

func (s Server) PostAccount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			s.log.WithError(err).Error("could not read request body")
			badRequest(w, "could not read request body")
			return
		}
		defer r.Body.Close()

		var account watdo.Account
		if err := json.Unmarshal(data, &account); err != nil {
			s.log.WithError(err).Error("could not marshal account")
			badRequest(w, "could not unmarshal account")
			return
		}

		id, err := s.service.SaveAccount(r.Context(), account)
		if err != nil {
			s.log.WithError(err).Error("could not save account")
			serverError(w, "could not save account")
			return
		}

		ok(w, []byte(id.String()))
	}
}

func (s Server) ListAccounts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accounts, err := s.service.ListAccounts(r.Context())
		if err != nil {
			s.log.WithError(err).Error("could not list accounts")
			serverError(w, "could not list accounts")
			return
		}

		data, err := json.Marshal(accounts)
		if err != nil {
			s.log.WithError(err).Error("could not marshal account listing")
			serverError(w, "could not list accounts")
			return
		}

		ok(w, data)
	}
}

func (s Server) GetAccount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		qAccountID := chi.URLParam(r, "accountID")
		log := s.log.WithField("accountID", qAccountID)
		accountID, err := uid.ParseString(qAccountID)
		if err != nil {
			log.WithError(err).Error("could not parse accountID")
			badRequest(w, "could not parse accountID")
			return
		}

		account, err := s.service.FetchAccount(r.Context(), accountID)
		if err != nil {
			log.WithError(err).Error("could not fetch account")
			serverError(w, "could not fetch account")
			return
		}

		data, err := json.Marshal(account)
		if err != nil {
			log.WithError(err).Error("could not marshal account")
			serverError(w, "could not marshal account")
		}

		ok(w, data)
	}
}
