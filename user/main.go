package main

import (
	"log"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.
		Methods("GET").
		Path("/users").
		HandlerFunc(getAll)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	users := []User {
		User{
			FirstName: "Firstname1",
			LastName: "Lastname1",
			UserName: "UserName1",
		},
		User{
			FirstName: "Firstname2",
			LastName: "Lastname2",
			UserName: "UserName2",
		},
	}

	j, _ := json.Marshal(users)

	w.Write(j)
}