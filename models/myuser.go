package models

import(
  "appengine"
  "appengine/datastore"
)

type User struct{
  Id int64
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

func setParentKey(c appengine.Context) *datastore.Key{
  return getParentKey(c, CATALOG_NAME, USER_NAME)
}

func (this *User) key(c appengine.Context) *datastore.Key{
  if this.Id == 0 {
    return datastore.NewIncompleteKey(c, USER_NAME, setParentKey(c))
  }
  return datastore.NewKey(c, USER_NAME, "", this.Id, setParentKey(c))
}

func (this *User) Save(c appengine.Context) (*datastore.Key, error){
  key := this.key(c)
  finalKey, err := datastore.Put(c,key,this)
  this.Id = finalKey.IntID()
  return finalKey, err
}

func GetUser(c appengine.Context, key *datastore.Key) (User, error){
  var user User
  err := datastore.Get(c, key, &user)
  user.Id = key.IntID()
  return user, err
}

func (this *User) Delete(c appengine.Context) (error){
  key := this.key(c)
  err := datastore.Delete(c, key)
  return err
}

func DeleteUserByKeyId(id int64, c appengine.Context) error{
  key := datastore.NewKey(c, USER_NAME, "", id, setParentKey(c))
  return datastore.Delete(c, key)
}

func ListUsers(c appengine.Context) ([]User, []*datastore.Key, error){
  query := datastore.NewQuery(USER_NAME).Ancestor(setParentKey(c))
  users := []User{}
  keys, err := query.GetAll(c, &users)
  return users, keys, err
}
