package main

import (
	"calendar_service/src/config"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	if err := config.Load(); err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(200)
		w.Write([]byte(`{"message":"welcome to calendar service"}`))
	})

	srv := &http.Server{Addr: config.Config.Port, Handler: r}
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
