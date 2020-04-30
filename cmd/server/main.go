package main

import (
	"github.com/eriktate/watdo/http"
	"github.com/eriktate/watdo/postgres"
	"github.com/eriktate/watdo/service"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	store, err := postgres.New(postgres.NewStoreOpts())
	if err != nil {
		logger.WithError(err).Fatal("could not connect to postgres")
	}

	service := service.NewAccountService(store)
	server := http.NewServer(
		http.WithLogger(logger),
		http.WithService(service),
	)

	if err := server.Listen(); err != nil {
		logger.WithError(err).Fatal("server crashed")
	}
}
