package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/necutya/card_validator/internal/config"
	"github.com/necutya/card_validator/internal/server/http"
	"github.com/necutya/card_validator/internal/server/http/controllers"
	cardservice "github.com/necutya/card_validator/internal/services/card"
	log "github.com/sirupsen/logrus"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go gracefulShutdown(cancel)

	cfg, err := config.New()
	if err != nil {
		log.WithError(err).Fatalf("cannot setup config")
	}

	cardService := cardservice.New()

	ctrl := controllers.New(cardService)
	router := http.NewRouter(ctrl)
	httpSrv := http.NewHTTPServer(router, cfg.HTTPPort)

	httpSrv.Run(ctx)
}

func gracefulShutdown(stop func()) {
	signalChannel := make(chan os.Signal, 1)

	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	signal.Notify(signalChannel, os.Interrupt, os.Interrupt)

	<-signalChannel
	stop()
}
