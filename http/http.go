package http

import (
	"fmt"
	"net/http"

	"github.com/eriktate/watdo"
	"github.com/eriktate/watdo/env"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

type ConfigOpt func(s Server) Server

type Server struct {
	host string
	port uint
	log  *logrus.Logger

	service watdo.WatdoService
}

func NewServer(opts ...ConfigOpt) Server {
	server := Server{
		host: env.GetString("WATDO_HOST", "localhost"),
		port: env.GetUint("WATDO_PORT", 8080),
		log:  logrus.New(),
	}

	for _, opt := range opts {
		server = opt(server)
	}

	return server
}

func WithHost(host string) ConfigOpt {
	return func(s Server) Server {
		s.host = host
		return s
	}
}

func WithPort(port uint) ConfigOpt {
	return func(s Server) Server {
		s.port = port
		return s
	}
}

func WithLogger(log *logrus.Logger) ConfigOpt {
	return func(s Server) Server {
		s.log = log
		return s
	}
}

func WithService(service watdo.WatdoService) ConfigOpt {
	return func(s Server) Server {
		s.service = service
		return s
	}
}

func (s Server) Listen() error {
	return http.ListenAndServe(fmt.Sprintf("%s:%d", s.host, s.port), s.buildRouter())
}

func (s Server) buildRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/account", s.PostAccount())
	r.Get("/account", s.ListAccounts())
	r.Get("/account/{accountID}", s.GetAccount())

	return r
}

func ok(w http.ResponseWriter, data []byte) {
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func noContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
	w.Write(nil)
}

func badRequest(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(msg))
}

func serverError(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(msg))
}
