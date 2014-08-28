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

const USER_NAME = "User"

func (this *User) key(c appengine.Context) *datastore.Key{
  if this.Id == 0 {
    return datastore.NewIncompleteKey(c, USER_NAME, nil)
  }
  return datastore.NewKey(c, USER_NAME, "", this.Id, nil)
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
  key := datastore.NewKey(c, USER_NAME, "", id, nil)
  return datastore.Delete(c, key)
}
