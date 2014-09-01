package models

import(
  "appengine"
  "appengine/datastore"
)

type User struct{
  Id string
  Name string
  Email string
}

type UserList struct{
  Users []User
}

const USER_NAME = "User"
const CATALOG_NAME = "Catalog"
const SCHEMA = "Schema"

func getParentKey(c appengine.Context, entityName string, schema string) *datastore.Key{
  return datastore.NewKey(c, entityName, SCHEMA, 0, nil)
}

func DeleteAllUsers(c appengine.Context) error {
  return datastore.RunInTransaction(c, func(c appengine.Context) error {
    ks, err := datastore.NewQuery(USER_NAME).KeysOnly().Ancestor(setParentKey(c)).GetAll(c, nil)
    if err != nil {
      return err
    }
    return datastore.DeleteMulti(c, ks)
  }, nil)
}

func setParentKey(c appengine.Context) *datastore.Key{
  return getParentKey(c, CATALOG_NAME, USER_NAME)
}

func (this *User) key(c appengine.Context) *datastore.Key{
  if this.Id == "" {
    return datastore.NewIncompleteKey(c, USER_NAME, setParentKey(c))
  }
  key, _ := datastore.DecodeKey(this.Id)
  return key
}

func (this *User) Save(c appengine.Context) (*datastore.Key, error){
  key := this.key(c)
  finalKey, err := datastore.Put(c,key,this)
  this.Id = finalKey.Encode()
  return finalKey, err
}

func GetUser(c appengine.Context, key *datastore.Key) (User, error){
  var user User
  err := datastore.Get(c, key, &user)
  user.Id = key.Encode()
  return user, err
}

func (this *User) Delete(c appengine.Context) (error){
  key := this.key(c)
  err := datastore.Delete(c, key)
  return err
}

func DeleteUserByKeyId(id string, c appengine.Context) error{
  key, _ := datastore.DecodeKey(id)
  return datastore.Delete(c, key)
}

func ListUsers(c appengine.Context) ([]User, []*datastore.Key, error){
  query := datastore.NewQuery(USER_NAME).Ancestor(setParentKey(c))
  users := make([]User,0,100)
  keys, err := query.GetAll(c, &users)
  for index, val := range keys{
    users[index].Id = val.Encode()
  }
  return users, keys, err
}
