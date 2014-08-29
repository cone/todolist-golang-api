### Go API example using:

 - GoConvey as a test suite
 - Google Appengine
 - Datastore
 - Go lang

#### Retrieve the User list
```
curl http://localhost:8080/api/users
```

#### Testing the POST request (Create a New User)

```
curl -H "Content-Type: application/json" -d '{ "Name": "User Name", "Email": "user@example.com"}' http://localhost:8080/api/users
```
