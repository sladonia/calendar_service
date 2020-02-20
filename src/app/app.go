package app

import (
	"calendar_service/src/config"
	"calendar_service/src/controllers"
	"calendar_service/src/datasources/postgres/calendardb"
	"calendar_service/src/logger"
	"calendar_service/src/middlewares/logging_middleware"
	"calendar_service/src/models"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/http"
)

func ConfigureApp() error {
	if err := config.Load(); err != nil {
		return err
	}
	if err := logger.InitLogger(config.Config.ServiceName, config.Config.LogLevel); err != nil {
		return err
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
		return err
	}
	calendardb.DB.LogMode(false)
	return nil
}

func InitApp() http.Handler {
	r := mux.NewRouter()
	r.NotFoundHandler = &controllers.NotFoundHandler{}
	r.HandleFunc("/", controllers.RootController.Get)

	r.HandleFunc("/user", controllers.UserController.Create).Methods("POST")
	r.HandleFunc("/user/{id}", controllers.UserController.Read).Methods("GET")
	r.HandleFunc("/user/{id}", controllers.UserController.Delete).Methods("DELETE")
	r.HandleFunc("/user/{id}", controllers.UserController.Update).Methods("POST")
	r.HandleFunc("/user/{user_id}/calendar", controllers.CalendarController.Create).Methods("POST")
	r.HandleFunc("/calendar/{calendar_id}", controllers.CalendarController.Read).Methods("GET")
	r.HandleFunc("/calendar/{calendar_id}", controllers.CalendarController.Update).Methods("POST")
	r.HandleFunc("/calendar/{calendar_id}", controllers.CalendarController.Delete).Methods("DELETE")
	r.HandleFunc("/calendar/{calendar_id}/appointment", controllers.AppointmentController.Create).Methods("POST")
	r.HandleFunc("/appointment/{appointment_id}", controllers.AppointmentController.Read).Methods("GET")
	r.HandleFunc("/appointment/{appointment_id}", controllers.AppointmentController.Update).Methods("POST")
	r.HandleFunc("/appointment/{appointment_id}", controllers.AppointmentController.Delete).Methods("DELETE")
	r.HandleFunc("/appointment/{appointment_id}/add-attendees", controllers.AppointmentController.AddAttendees).Methods("POST")
	r.HandleFunc("/appointment/{appointment_id}/remove-attendees", controllers.AppointmentController.RemoveAttendees).Methods("POST")

	r.Use(logging_middlewaer.LoggingMw)

	return r
}
