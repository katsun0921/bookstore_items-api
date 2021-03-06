package app

import (
  "github.com/katsun0921/bookstore_items-api/src/controllers"
	"net/http"
)

func mapUrls() {
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)
	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
  router.HandleFunc("/items/{id}", controllers.ItemsController.Get).Methods(http.MethodPost)
  router.HandleFunc("/items/search", controllers.ItemsController.Search).Methods(http.MethodPost)
}
