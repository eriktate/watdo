package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/eriktate/wrkhub"
	"github.com/eriktate/wrkhub/uid"
	"github.com/go-chi/chi"
)

func (s Server) PostProject() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			s.log.WithError(err).Error("could not read request body")
			badRequest(w, "could not read request body")
			return
		}
		defer r.Body.Close()

		var project wrkhub.Project
		if err := json.Unmarshal(data, &project); err != nil {
			s.log.WithError(err).Error("could not marshal project")
			badRequest(w, "could not unmarshal project")
			return
		}

		id, err := s.service.SaveProject(r.Context(), project)
		if err != nil {
			s.log.WithError(err).Error("could not save project")
			serverError(w, "could not save project")
			return
		}

		ok(w, []byte(id.JSONString()))
	}
}

func (s Server) ListProjects() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accountID, err := uid.ParseString(r.URL.Query().Get("accountId"))
		if err != nil {
			s.log.WithError(err).WithField("accountID", accountID).Error("could not parse accountID")
			badRequest(w, "could not parse accountID")
			return
		}

		req := wrkhub.ListProjectsReq{AccountID: accountID}
		log := s.log.WithField("req", req)
		projects, err := s.service.ListProjects(r.Context(), req)
		if err != nil {
			log.WithError(err).Error("could not list projects")
			serverError(w, "could not list projects")
			return
		}

		data, err := json.Marshal(projects)
		if err != nil {
			log.WithError(err).Error("could not marshal project listing")
			serverError(w, "could not list projects")
			return
		}

		ok(w, data)
	}
}

func (s Server) GetProject() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		qProjectID := chi.URLParam(r, "projectID")
		log := s.log.WithField("projectID", qProjectID)
		projectID, err := uid.ParseString(qProjectID)
		if err != nil {
			log.WithError(err).Error("could not parse projectID")
			badRequest(w, "could not parse projectID")
			return
		}

		project, err := s.service.FetchProject(r.Context(), projectID)
		if err != nil {
			log.WithError(err).Error("could not fetch project")
			respondWithError(w, err)
			return
		}

		data, err := json.Marshal(project)
		if err != nil {
			log.WithError(err).Error("could not marshal project")
			serverError(w, "could not marshal project")
		}

		ok(w, data)
	}
}
