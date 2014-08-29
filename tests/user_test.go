package tests

import(
  . "todolistapi/models"
  "testing"
  . "github.com/smartystreets/goconvey/convey"
  utils "todolistapi/utils"
)

func TestUserListRequest(t *testing.T) {
  c := utils.InitializeContext()
  defer c.Close()

  Convey("Given a list of users created", t, func(){
    utils.PopulateUserList(c, 100)
    Convey("When I get all the users created", func(){
      users, _, _ := ListUsers(c)
      Convey("Then the quantity of user should be equal to 5000", func(){
        So(len(users), ShouldEqual, 100)
      })
    })
  })
}

func TestUserCreation(t *testing.T){

  c := utils.InitializeContext() 
  defer c.Close()
  Convey("Given a user is created", t, func(){
    user := User{ 0, "Carlos", "coneramu@gmail.com"}
    userKey, _ := user.Save(c)
    Convey("When the user it is retrieved", func(){
      userFromDB, _ := GetUser(c, userKey)
      wantedName := "Carlos"
      Convey("Then the user name should be " + wantedName, func(){
        So(userFromDB.Name, ShouldEqual, wantedName)
      })
    })
  })
}

func TestUserDeletion(t *testing.T) {
  c := utils.InitializeContext()
  defer c.Close()

  Convey("Given a user is created", t, func(){
    user := User{ 0, "Carlos", "coneramu@gmail.com"}
    userKey, _ := user.Save(c)

    Convey("When the user is deleted using the userKey", func(){
      user.Delete(c)
      _, err := GetUser(c, userKey)
      Convey("Then the user should not exist in the database", func(){
        So(err.Error(), ShouldEqual, "datastore: no such entity")
      })
    })

    Convey("When the user is deleted using the Id", func(){
      DeleteUserByKeyId(user.Id, c)

      _, err := GetUser(c, userKey)
      Convey("Then the user should not exist in the database", func(){
        So(err.Error(), ShouldEqual, "datastore: no such entity")
      })
    })

  })
}

func TestUserUpdation(t *testing.T){
  c := utils.InitializeContext()
  defer c.Close()

  Convey("Given a user is created", t, func(){
    user := User{ 0, "Carlos", "coneramu@gmail.com"}
    user.Save(c)

   Convey("When the user is updated", func(){
     newEmail := "carlos.gutierrez@crowdint.com"
     user.Email = newEmail
     userKey, _ := user.Save(c)
     userFromBd, _ := GetUser(c, userKey)
     Convey("The user email should be '" + newEmail + "'", func(){
      So(userFromBd.Email, ShouldEqual, newEmail)
     })
   })
  })
}

