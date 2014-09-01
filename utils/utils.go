package utils

import(
  "appengine"
  "appengine/aetest"
  "encoding/json"
  "io"
  "todolistapi/models"
)

func InitializeContext() appengine.Context{
  c := appengine.NewContext(nil)
  return c
}

func InitializeTestContext() aetest.Context{
  c, _ := aetest.NewContext(nil)
  return c
}

func PopulateUserList(c appengine.Context, quantity int) {
  for i:= 0; i < quantity; i++ {
    email := "user+" + string(i) + "@example.com"
    name := "user " + string(i) + " name"
    user := models.User{
      "",
      name,
      email,
    }
    user.Save(c)
  }
}


func DecodeUser(r io.ReadCloser)(*models.User, error){
  defer r.Close()
  var user models.User
  err := json.NewDecoder(r).Decode(&user)
  return &user, err
}

func DecodeUserList(r io.ReadCloser)(*models.UserList, error){
  defer r.Close()
  var userlist models.UserList
  err := json.NewDecoder(r).Decode(&userlist)
  return &userlist, err
}

