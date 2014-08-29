package controllers

import (
  "appengine"
  "todolistapi/models"
  "encoding/json"
)

func IndexUser(c appengine.Context) string{
  users, _, _:= models.ListUsers(c)
  var userList models.UserList
  userList.Users = users
  //json.NewEncoder(w).Encode(userList)
  jsonString, _ := json.Marshal(userList)
  return string(jsonString)
}

//func CreateUser(){

//}

//func UpdateUser(){

//}

//func DeleteUser(){

//}

//func GetUser(){

//}
