package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"ordinary-system/user/data"
	"ordinary-system/user/model"
	"ordinary-system/utility"

	"github.com/gorilla/mux"
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
	users, err := data.GetUsers()

	if err != nil {
		utility.WriteInternalError(w, err)
		return
	}

	utility.WriteResponse(w, users)
}

func create(w http.ResponseWriter, r *http.Request) {
	user, err := getUserFromBody(r.Body)
	if err != nil {
		utility.WriteBadRequest(w, err)
		return
	}

	user, err = data.CreateUser(user)
	if err != nil {
		utility.WriteInternalError(w, err)
		return
	}

	utility.WriteResponse(w, user)
}

func getUserFromBody(body io.Reader) (model.User, error) {
	rawBody, _ := ioutil.ReadAll(body)
	var u model.User
	if err := json.Unmarshal(rawBody, &u); err != nil {
		return model.User{}, err
	}
	return u, nil
}
