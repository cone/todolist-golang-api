package main

import (
  "net/http"
  "todolistapi/routes"
)

func init(){
  m := routes.Router()
  http.Handle("/", m)
}

