package items

import (
  "encoding/json"
  "errors"
  "fmt"
  "github.com/katsun0921/bookstore_items-api/clients/elasticsearch"
	"github.com/katsun0921/bookstore_utils-go/rest_errors"
  "strings"
)

const (
	indexItems = "items"
	typeItem   = "item"
)

func (i *Item) Save() rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, typeItem, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when tyring to save", errors.New("database error"))
	}
	i.Id = result.Id
	return nil
}

func (i *Item) Get() rest_errors.RestErr {
  itemId := i.Id
  result, err := elasticsearch.Client.Get(indexItems, typeItem, i.Id)
  if err != nil {
    if strings.Contains(err.Error(), "404") {
      return rest_errors.NewNotFoundError(fmt.Sprintf("no item found with id %s", i.Id), errors.New("database error"))
    }
    return rest_errors.NewInternalServerError(fmt.Sprintf("error when trying to get id %s", i.Id), errors.New("database error"))
  }
  bytes, err := result.Source.MarshalJSON()
  if err != nil {
    return rest_errors.NewInternalServerError("error when trying to parse database response", errors.New("database error"))
  }
  if err != json.Unmarshal(bytes, &i) {
    return rest_errors.NewInternalServerError("error when trying to parse database response", errors.New("database error"))
  }
  i.Id = itemId
  return nil
}