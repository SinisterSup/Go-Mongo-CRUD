package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SinisterSup/Go-Mongo-CRUD/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
  session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
  return &UserController{s}
}

func (uc UserController) GetUser (w http.ResponseWriter, r *http.Request, p httprouter.Params) {
  id := p.ByName("id")
  if !bson.IsObjectIdHex(id) {
    w.WriteHeader(http.StatusNotFound)
  }

  oid := bson.ObjectIdHex(id)

  u := models.User{}

  if err := uc.session.DB("mongo-golang").C("users").FindId(oid).One(&u); err != nil {
    w.WriteHeader(http.StatusNotFound)
    return
  }

  uj, err := json.Marshal(u)
  if err != nil {
    fmt.Println(err)
    panic(err)
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  u := models.User{}

  err := json.NewDecoder(r.Body).Decode(&u)
  if err != nil {
    panic(err)
  }

  u.ID = bson.NewObjectId()

  uc.session.DB("mongo-golang").C("users").Insert(u)

  uj, err := json.Marshal(u)

  if err != nil {
    fmt.Println(err)
    panic(err)
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
  fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
  id := p.ByName("id")
  if !bson.IsObjectIdHex(id) {
    w.WriteHeader(http.StatusNotFound)
    return
  }

  oid := bson.IsObjectIdHex(id)

  if err := uc.session.DB("mongo-golang").C("users").RemoveId(oid); err != nil {
    w.WriteHeader(http.StatusNotFound)
  }

  w.WriteHeader(http.StatusOK)
  fmt.Fprintf(w, "Deleterd user %v\n", oid)
}
