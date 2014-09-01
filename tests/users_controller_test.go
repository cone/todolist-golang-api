package tests

import(
  "testing"
  . "github.com/smartystreets/goconvey/convey"
  "net/http"
  "log"
  "bytes"
  "net/http/httptest"
  "todolistapi/controllers"
  "todolistapi/utils"
  "io/ioutil"
)


func TestJsonUserList(t *testing.T) {
  c := utils.InitializeTestContext()

  defer c.Close()

  ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    controllers.DoIndexUser(c, w, r)
  }))

  defer ts.Close()

  Convey("Given a userList is created", t, func(){

    Convey("When an userList is created", func(){

      utils.PopulateUserList(c, 4)

      res, err := http.Get(ts.URL)

      if err != nil {
        log.Fatal(err)
      }

      userList, _ := utils.DecodeUserList(res.Body)
      Convey("Then the number of users created should be 4", func(){
        So(len(userList.Users), ShouldEqual, 4 )
      })
    })
  })
}

func TestCreateUserEndpoint(t *testing.T) {
  c := utils.InitializeTestContext()
  defer c.Close()

  stringParams :=  `{ "Name": "User Name", "Email": "user@example.com"}`
  reader := bytes.NewReader([]byte(stringParams))

  ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    controllers.DoCreateUser(c, w, r)
  }))

  defer ts.Close()

  Convey("Given a user is created", t, func(){
    r, _ := http.NewRequest("POST", ts.URL, reader)
    client := new(http.Client)
    response, _ := client.Do(r)
    str, _ := ioutil.ReadAll(response.Body)
    Convey("The user name should be: User Name", func(){
      So(string(str), ShouldEqual, "User created succesfuly with key")
    })

    Convey("The response status should be 200", func(){
      So(response.Status, ShouldEqual, "200 OK")
    })

  })
}


