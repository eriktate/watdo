package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/eriktate/wrkhub"
	"github.com/eriktate/wrkhub/uid"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

func PostTask(service wrkhub.TaskService, log *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.WithError(err).Error("could not read request body")
			badRequest(w, "could not read request body")
			return
		}
		defer r.Body.Close()

		var task wrkhub.Task
		if err := json.Unmarshal(data, &task); err != nil {
			log.WithError(err).Error("could not marshal task")
			badRequest(w, "could not unmarshal task")
			return
		}

		id, err := service.SaveTask(r.Context(), task)
		if err != nil {
			log.WithError(err).Error("could not save task")
			serverError(w, "could not save task")
			return
		}

		ok(w, []byte(id.JSONString()))
	}
}

func ListTasks(service wrkhub.TaskService, log *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		projectID, err := uid.ParseString(r.URL.Query().Get("projectId"))
		if err != nil {
			log.WithError(err).WithField("projectID", projectID).Error("could not parse projectID")
			badRequest(w, "could not parse projectID")
			return
		}

		req := wrkhub.ListTasksReq{ProjectID: projectID}
		log := log.WithField("req", req)
		tasks, err := service.ListTasks(r.Context(), req)
		if err != nil {
			log.WithError(err).Error("could not list tasks")
			serverError(w, "could not list tasks")
			return
		}

		data, err := json.Marshal(tasks)
		if err != nil {
			log.WithError(err).Error("could not marshal task listing")
			serverError(w, "could not list tasks")
			return
		}

		ok(w, data)
	}
}

func GetTask(service wrkhub.TaskService, log *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		qTaskID := chi.URLParam(r, "taskID")
		log := log.WithField("taskID", qTaskID)
		taskID, err := uid.ParseString(qTaskID)
		if err != nil {
			log.WithError(err).Error("could not parse taskID")
			badRequest(w, "could not parse taskID")
			return
		}

		task, err := service.FetchTask(r.Context(), taskID)
		if err != nil {
			log.WithError(err).Error("could not fetch task")
			serverError(w, "could not fetch task")
			return
		}

		data, err := json.Marshal(task)
		if err != nil {
			log.WithError(err).Error("could not marshal task")
			serverError(w, "could not marshal task")
		}

		ok(w, data)
	}
}
