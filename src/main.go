package main

import (
  "github.com/katsun0921/bookstore_items-api/src/app"
  "os"
)

func main() {
  os.Setenv("LOG_LEVEL", "info")
  app.StartApplication()
}
