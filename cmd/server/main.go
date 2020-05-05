package main

import (
	"github.com/eriktate/wrkhub/http"
	"github.com/eriktate/wrkhub/postgres"
	"github.com/eriktate/wrkhub/service"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	store, err := postgres.New()
	if err != nil {
		logger.WithError(err).Fatal("could not connect to postgres")
	}

	service := service.NewService(store, store, store)
	server := http.NewServer(
		http.WithLogger(logger),
		http.WithService(service),
	)

	if err := server.Listen(); err != nil {
		logger.WithError(err).Fatal("server crashed")
	}
}
