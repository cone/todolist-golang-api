package main

import (
  "appengine"
  "net/http"
  "fmt"
  "io"
  "todolistapi/controllers"
  "todolistapi/models"
  "encoding/json"
  "github.com/gorilla/mux"
)

func init(){
  m := mux.NewRouter()

  m.HandleFunc("/api/users", handleUsers).Methods("GET")
  m.HandleFunc("/api/users", createHdlr).Methods("POST")
  m.HandleFunc("/api/users", deleteUsers).Methods("DELETE")

  m.HandleFunc("/api/users/{key}", deleteUserHandler).Methods("DELETE")
  m.HandleFunc("/api/users/{key}", updateHdlr).Methods("PUT")
  //m.HandleFunc("/api/users/{key}", showUserDetails).Methods("GET")
  http.Handle("/", m)
}

func updateHdlr(w http.ResponseWriter, r *http.Request){
  c := appengine.NewContext(r)
  user, _ := decodeUser(r.Body)

  vars := mux.Vars(r)
  key := vars["key"] 

  user.Id = key
  _, err := user.Save(c)

  if err != nil {
    fmt.Fprintf(w, "%s %s", "the User wasn't found:", err.Error())
  } else {
    fmt.Fprintf(w, "%s", "the User was updated succesfuly")
  }
}


func deleteUsers(w http.ResponseWriter, r *http.Request){
  c := appengine.NewContext(r)
  models.DeleteAllUsers(c)
  fmt.Fprintf(w, "%v", "Successfully")
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
  //c := appengine.NewContext(r)
}

func handleUsers(w http.ResponseWriter, r *http.Request){
  c := appengine.NewContext(r)
  jsonString := controllers.IndexUser(c)
  fmt.Fprintf(w, "%v", jsonString)
}

func createHdlr(w http.ResponseWriter, r *http.Request){
  c := appengine.NewContext(r)
  user, err := decodeUser(r.Body)
  controllers.DoCreateUser(c, user)
  if err != nil {
    fmt.Fprintf(w, "%s %s", "User was not created successfully", err.Error())
  } else {
    fmt.Fprintf(w, "%s %s", "User created succesfuly with key: ", user.Id)
  }
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request ) {
  c := appengine.NewContext(r)

  vars := mux.Vars(r)
  key := vars["key"] 

  err := controllers.DeleteUser(c, key)
  message := "The user deletion was"
  if err != nil {
    fmt.Fprintf(w, "%s %s", message, "Unsuccessful")
  }else{
    fmt.Fprintf(w, "%s %s", message, "Successful")
  }
}

func decodeUser(r io.ReadCloser)(*models.User, error){
  defer r.Close()
  var user models.User
  err := json.NewDecoder(r).Decode(&user)
  return &user, err
}
