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
	initResponse(w)
	
	users, err := data.GetUsers()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func create(w http.ResponseWriter, r *http.Request) {
	initResponse(w)

	user, err := getUserFromBody(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	user, err = data.CreateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func initResponse (w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func getUserFromBody(body io.Reader) (model.User, error) {
	rawBody, _ := ioutil.ReadAll(body)
	var u model.User
	if err := json.Unmarshal(rawBody, &u); err != nil {
		return model.User{}, err
	}
	return u, nil
}