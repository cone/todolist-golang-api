package controllers

import (
  "appengine"
  "net/http"
  "fmt"
  "todolistapi/models"
  "encoding/json"
  "github.com/gorilla/mux"
  "todolistapi/utils"
)

func IndexUser(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)
  DoIndexUser(c, w, r)
}

func DoIndexUser(c appengine.Context, w http.ResponseWriter, r *http.Request) {
  users, _, _:= models.ListUsers(c)
  var userList models.UserList
  userList.Users = users
  jsonString, _ := json.Marshal(userList)
  fmt.Fprintf(w, "%v", string(jsonString))
}

func DeleteUser(c appengine.Context, userId string) error{
  return models.DeleteUserByKeyId(userId, c)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)
  DoCreateUser(c, w, r)
}

func DoCreateUser(c appengine.Context, w http.ResponseWriter, r *http.Request) {
  user, _ := utils.DecodeUser(r.Body)
  _, err := user.Save(c)
  if err != nil {
    fmt.Fprintf(w, "%s %s", "User was not created successfully", err.Error())
  } else {
    fmt.Fprintf(w, "%s", "User created succesfuly with key")
  }
}

func DeleteUsers(w http.ResponseWriter, r *http.Request){
  c := appengine.NewContext(r)
  models.DeleteAllUsers(c)
  fmt.Fprintf(w, "%v", "Successfully")
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request ) {
  c := appengine.NewContext(r)

  vars := mux.Vars(r)
  key := vars["key"] 

  err := DeleteUser(c, key)
  message := "The user deletion was"
  if err != nil {
    fmt.Fprintf(w, "%s %s", message, "Unsuccessful")
  }else{
    fmt.Fprintf(w, "%s %s", message, "Successful")
  }
}

func UpdateHdlr(w http.ResponseWriter, r *http.Request){
  c := appengine.NewContext(r)
  user, _ := utils.DecodeUser(r.Body)

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

