package main

import (
  "appengine"
  "net/http"
  "fmt"
  "io"
  "todolistapi/controllers"
  "todolistapi/models"
  "encoding/json"
)

func init(){
  http.HandleFunc("/api/users", handler)
}

func handler(w http.ResponseWriter, r *http.Request){
  c := appengine.NewContext(r)
  handleUsers(c, r, w)
}

func handleUsers(c appengine.Context, r *http.Request, w http.ResponseWriter){

  switch r.Method  {
  case "POST" :
    createHdlr(w, r)
  case "GET"  :
    jsonString := controllers.IndexUser(c)
    fmt.Fprintf(w, "%v", jsonString)
  case "DELETE" :
    //deleteUserHandler(w, r)
    fmt.Fprintf(w, "%v", "DELETE")
  }
}

func createHdlr(w http.ResponseWriter, r *http.Request){
  c := appengine.NewContext(r)
  user, err := decodeUser(r.Body)
  fmt.Fprintf(w, "%s", "body", r.Body)
  controllers.DoCreateUser(c, user)
  if err != nil {
    fmt.Fprintf(w, "%s %s", "User was not created successfully", err.Error())
  } else {
    fmt.Fprintf(w, "%s", "User created succesfuly")
  }
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request ) {
  c := appengine.NewContext(r)
  user, _ := decodeUser(r.Body)
  err := controllers.DeleteUser(c, user.Id)
  message := "The user deletion was"
  if err != nil {
    fmt.Fprintf(w, "%s %s", message, "Unsuccessful")
  }else{
    fmt.Fprintf(w, "%s %s", message, "Successful")
  }
  //params := r.URL.Query()
  //fmt.Fprintf(w, "%v", params)
}

func decodeUser(r io.ReadCloser)(*models.User, error){
  defer r.Close()
  var user models.User
  err := json.NewDecoder(r).Decode(&user)
  return &user, err
}
