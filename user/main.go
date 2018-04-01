package main

import (
	"io/ioutil"
	"io"
	"log"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"ordinary-system/user/model"
	"ordinary-system/user/data"
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
	
	users, err := data.GetUsers()

	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	
	json.NewEncoder(w).Encode(users)
}

func create(w http.ResponseWriter, r *http.Request) {
	initOKResponse(w)

	user := getUserFromBody(r.Body)

	user, err := data.CreateUser(user)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func initOKResponse (w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func getUserFromBody(body io.Reader) model.User {
	rawBody, _ := ioutil.ReadAll(body)
	var u model.User
	if err := json.Unmarshal(rawBody, &u); err != nil {
		panic(err)
	}
	return u
}