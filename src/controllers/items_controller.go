package controllers

import (
  "encoding/json"
  "github.com/gorilla/mux"
  "github.com/katsun0921/bookstore_items-api/src/domain/items"
  "github.com/katsun0921/bookstore_items-api/src/domain/queries"
  "github.com/katsun0921/bookstore_items-api/src/services"
  "github.com/katsun0921/bookstore_items-api/src/utils/http_utils"
  "github.com/katsun0921/bookstore_oauth-go/oauth"
  "github.com/katsun0921/bookstore_utils-go/rest_errors"
  "io/ioutil"
  "net/http"
  "strings"
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
  Search(w http.ResponseWriter, r *http.Request)
}

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsController struct{}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		//TODO: Return error to the user.
    // http_utils.ResponseError(w, err)
		return
	}

	sellerId := oauth.GetClientId(r)
	if sellerId == 0 {
	  restErr := rest_errors.NewBadRequestError("invalid request body", nil)
    http_utils.ResponseError(w, restErr)
	  return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
	  restErr := rest_errors.NewBadRequestError("invalid request body", err)
    http_utils.ResponseError(w, restErr)
	  return
  }
  defer r.Body.Close()

  var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
    restErr := rest_errors.NewBadRequestError("invalid item json body", err)
    http_utils.ResponseError(w, restErr)
  }

  itemRequest.Seller = oauth.GetCallerId(r)


	result, createErr := services.ItemService.Create(itemRequest)
	if createErr != nil {
		//TODO: Return error json to the user.
	  http_utils.ResponseError(w, createErr)
	  return
	}
	//TODO: Return created item as json with HTTP status 201 - Created
  http_utils.ResponseJson(w, http.StatusCreated, result)

}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  itemId := strings.TrimSpace(vars["id"])

  item, err := services.ItemService.Get(itemId)
  if err != nil {
    http_utils.ResponseError(w, err)
    return
  }
  http_utils.ResponseJson(w, http.StatusOK, item)
}

func (c *itemsController) Search(w http.ResponseWriter, r *http.Request) {
  bytes, err := ioutil.ReadAll(r.Body)
  if err != nil {
    apiErr := rest_errors.NewBadRequestError("invalid json body", err)
    http_utils.ResponseError(w, apiErr)
    return
  }
  defer r.Body.Close()

  var query queries.EsQuery
  if err := json.Unmarshal(bytes, &query); err != nil {
    apiErr := rest_errors.NewBadRequestError("invalid json body", err)
    http_utils.ResponseError(w, apiErr)
    return
  }

  searchItems, searchErr := services.ItemService.Search(query)
  if err != nil {
    http_utils.ResponseError(w, searchErr)
    return
  }
  http_utils.ResponseJson(w, http.StatusOK, searchItems)
}
