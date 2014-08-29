### Go API example using:

 - GoConvey as a test suite
 - Google Appengine
 - Datastore
 - Go lang

#### Retrieve the User list
```
curl http://localhost:8080/api/users
```

### Create request using CURL

#### (Creating a New User)

```
curl -H "Content-Type: application/json" -d '{ "Name": "User Name", "Email": "user@example.com"}' http://localhost:8080/api/users
```

#### Update a existent User
```
curl -X PUT -d '{"Name": "New user name"}' http://localhost:8080/api/users/ag5kZXZ-dG9kb2xpc3RjaXIkCxIHQ2F0YWxvZyIGU2NoZW1hDAsSBFVzZXIYgICAgIDg1wgM
```

#### Delete all the Users

```
curl -X DELETE http://localhost:8080/api/users
```
#### Delete an specific User
```
curl -X DELETE http://localhost:8080/api/users/ag5kZXZ-dG9kb2xpc3RjaXIkCxIHQ2F0YWxvZyIGU2NoZW1hDAsSBFVzZXIYgICAgIDglwoM
```
