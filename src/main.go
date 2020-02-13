package main

import (
	"calendar_service/src/config"
	"calendar_service/src/logger"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := config.Load(); err != nil {
		panic(err)
	}
	if err := logger.InitLogger(config.Config.ServiceName, config.Config.LogLevel); err != nil {
		panic(err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(200)
		w.Write([]byte(`{"message":"welcome to calendar service"}`))
	})

	srv := &http.Server{Addr: config.Config.Port, Handler: r}
	logger.Logger.Infof("start listening on port %s", config.Config.Port)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	<-done
	logger.Logger.Info("shutting down gracefully")
	logger.Logger.Sync()
}
