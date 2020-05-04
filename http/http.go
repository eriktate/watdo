package http

import (
	"fmt"
	"net/http"

	"github.com/eriktate/wrkhub"
	"github.com/eriktate/wrkhub/env"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/sirupsen/logrus"
)

type ConfigOpt func(s Server) Server

type Server struct {
	host string
	port uint
	log  *logrus.Logger

	service wrkhub.WrkhubService
}

func NewServer(opts ...ConfigOpt) Server {
	server := Server{
		host: env.GetString("WRKHUB_HOST", "localhost"),
		port: env.GetUint("WRKHUB_PORT", 8080),
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

func WithService(service wrkhub.WrkhubService) ConfigOpt {
	return func(s Server) Server {
		s.service = service
		return s
	}
}

func (s Server) Listen() error {
	s.log.WithField("host", s.host).WithField("port", s.port).Info("starting server...")
	return http.ListenAndServe(fmt.Sprintf("%s:%d", s.host, s.port), s.buildRouter())
}

func (s Server) buildRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))

	r.Post("/account", PostAccount(s.service, s.log))
	r.Get("/account", s.ListAccounts())
	r.Get("/account/{accountID}", s.GetAccount())

	r.Post("/project", s.PostProject())
	r.Get("/project", s.ListProjects())
	r.Get("/project/{projectID}", s.GetProject())

	return r
}

func ok(w http.ResponseWriter, data []byte) {
	w.Header().Add("Content-Type", "application/json")
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

func respondWithError(w http.ResponseWriter, err error) {
	werr, ok := err.(wrkhub.WrkhubErr)
	if !ok {
		serverError(w, "something went wrong")
		return
	}

	safe := werr.SafeMsg()
	switch werr.ErrType() {
	case wrkhub.ErrInvalid:
		badRequest(w, safe)
	default:
		serverError(w, safe)
	}
}
