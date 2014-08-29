package tests

import(
  "testing"
  . "github.com/smartystreets/goconvey/convey"
  //"todolistapi/models"
  "todolistapi/utils"
  //"encoding/json"
  "todolistapi/controllers"
)

func Test(t *testing.T) {
  c := utils.InitializeContext()

  defer c.Close()

  //Convey("")
  
}

func TestJsonUserList(t *testing.T) {
  c := utils.InitializeContext()

  defer c.Close()

  Convey("Given a userList is created", t, func(){

    utils.PopulateUserList(c, 2)
    Convey("When the userList is json marshalled", func(){
      jsonStringTest := `{"Users":[{"Id":0,"Name":"user \u0001 name","Email":"user+\u0001@example.com"},{"Id":0,"Name":"user \u0000 name","Email":"user+\u0000@example.com"}]}`
      jsonString := controllers.IndexUser(c)

      Convey("Then the userList should be in json format", func(){
        So(jsonStringTest, ShouldEqual, string(jsonString))
      })
    })
  })
}

