package main

import (
	"calendar_service/src/app"
	"calendar_service/src/config"
	"calendar_service/src/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := app.ConfigureApp(); err != nil {
		panic(err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	r := app.InitApp()

	logger.Logger.Infof("start listening on port %s", config.Config.Port)
	go func() {
		if err := http.ListenAndServe(":8080", r); err != nil {
			logger.Logger.Fatalw("unable to start the server", "error", err)
		}
	}()

	<-done
	logger.Logger.Info("shutting down gracefully")
	logger.Logger.Sync()
}
