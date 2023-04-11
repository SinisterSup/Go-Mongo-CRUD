package main

import (
  "fmt"
  "os"
  "github.com/julienschmidt/httprouter"
  "gopkg.in/mgo.v2"
  "net/http"
  "github.com/SinisterSup/Go-Mongo-CRUD/controllers"

)

func main() {
  r := httprouter.New()
  uc := controllers.NewUserController(getSession())
  r.GET("/user/:id", uc.GetUser)
  r.POST("/user", uc.CreateUser)
  r.DELETE("/user/:id", uc.DeleteUser)
  err := http.ListenAndServe("localhost:8080", r)
  if err != nil {
    fmt.Println("Failed to start server:", err)
    os.Exit(1)
  }
  
}

func getSession() *mgo.Session {
  s, err := mgo.Dial("mongodb://localhost:27017")
  if err != nil {
    fmt.Println("Failed to connect to MongoDB:", err)
    os.Exit(1)
  }
  return s
}