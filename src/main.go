package main

import (
	"calendar_service/src/config"
	"calendar_service/src/controllers"
	"calendar_service/src/datasources/postgres/calendardb"
	"calendar_service/src/logger"
	logging_middlewaer "calendar_service/src/middlewares/logging_middleware"
	"calendar_service/src/models"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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

	var err error
	calendardb.DB, err = models.InitDbConnection(
		config.Config.CalendarDb.User,
		config.Config.CalendarDb.Password,
		config.Config.CalendarDb.DbName,
		config.Config.CalendarDb.SslMode,
		config.Config.CalendarDb.MaxOpenConnections,
		config.Config.CalendarDb.MaxIdleConnections,
		config.Config.CalendarDb.ConnectionMaxLifetime,
	)
	if err != nil {
		logger.Logger.Fatalw("unable to establish db connection", "error", err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	r := mux.NewRouter()
	r.NotFoundHandler = &controllers.NotFoundHandler{}
	r.HandleFunc("/", controllers.RootController.Get)

	r.Use(logging_middlewaer.LoggingMw)

	srv := &http.Server{Addr: config.Config.Port, Handler: r}
	logger.Logger.Infof("start listening on port %s", config.Config.Port)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Logger.Fatalw("unable to start the server", "error", err)
		}
	}()

	<-done
	logger.Logger.Info("shutting down gracefully")
	logger.Logger.Sync()
}
