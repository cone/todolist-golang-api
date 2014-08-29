package utils

import(
  "appengine"
  "todolistapi/models"
  "appengine/aetest"
)

func InitializeContext() aetest.Context{
  c, _ := aetest.NewContext(nil)
  return c
}

func PopulateUserList(c appengine.Context, quantity int) {
  for i:= 0; i < quantity; i++ {
    email := "user+" + string(i) + "@example.com"
    name := "user " + string(i) + " name"
    user := models.User{
      0,
      name,
      email,
    }
    user.Save(c)
  }
}

