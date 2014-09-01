package routes

import(
  "github.com/gorilla/mux"
  "todolistapi/controllers"
)

func Router() *mux.Router{
  m := mux.NewRouter()

  m.HandleFunc("/api/users", controllers.IndexUser).Methods("GET")
  m.HandleFunc("/api/users", controllers.CreateUser).Methods("POST")
  m.HandleFunc("/api/users", controllers.DeleteUsers).Methods("DELETE")

  m.HandleFunc("/api/users/{key}", controllers.DeleteUserHandler).Methods("DELETE")
  m.HandleFunc("/api/users/{key}", controllers.UpdateHdlr).Methods("PUT")

  return m
}

