package app

import (
	"github.com/gorilla/mux"
  "github.com/katsun0921/bookstore_items-api/clients/elasticsearch"
  "net/http"
	"time"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
  elasticsearch.Init()

	mapUrls()

	srv := &http.Server{
		Handler:      router,
		Addr:         "localhost:8080",
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
