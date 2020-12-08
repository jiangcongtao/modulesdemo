package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	mtx := mux.NewRouter()
	mtx.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.WithField("user-agent", r.UserAgent()).Info("accessed home page")
		_, _ = w.Write([]byte("Hello World!"))
	})

	srv := &http.Server{
		Handler:      mtx,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}