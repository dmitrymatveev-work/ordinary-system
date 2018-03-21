package main

import (
	"io/ioutil"
	"io"
	"log"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.
		Methods("GET").
		Path("/api/users").
		HandlerFunc(getAll)

	router.
		Methods("POST").
		Path("/api/users").
		HandlerFunc(create)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getAll(w http.ResponseWriter, r *http.Request) {
	initOKResponse(w)
	
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("users").C("users")
	var users []User
	err = c.Find(bson.M{}).All(&users)

	json.NewEncoder(w).Encode(users)
}

func create(w http.ResponseWriter, r *http.Request) {
	initOKResponse(w)

	user := getUserFromBody(r.Body)

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("users").C("users")
	err = c.Insert(&user)
	if err != nil {
		log.Fatal(err)
	}
}

func initOKResponse (w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func getUserFromBody(body io.Reader) User {
	rawBody, _ := ioutil.ReadAll(body)
	var user User
	if err := json.Unmarshal(rawBody, &user); err != nil {
		panic(err)
	}
	return user
}