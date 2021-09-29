package controllers

import (
  "fmt"
  "github.com/katsun0921/bookstore_items-api/domain/items"
  "github.com/katsun0921/bookstore_items-api/services"
  "github.com/katsun0921/bookstore_oauth-go/oauth"
  "net/http"
)

type itemsControllerInterface interface {
  Create(w http.ResponseWriter, r *http.Request)
  Get(w http.ResponseWriter, r *http.Request)
}

var(
  ItemsController itemsControllerInterface = &itemsController{}
)

type itemsController struct {}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
  if err := oauth.AuthenticateRequest(r); err != nil {
    //TODO: Return error to the user.
    return
  }

  item := items.Item{
    Seller: oauth.GetCallerId(r),
  }

  result, err := services.ItemService.Create(item)
  if err != nil {
    //TODO: Return error json to the user.
  }
  fmt.Println(result)
  //TODO: Return created item as json with HTTP status 201 - Created
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {}
