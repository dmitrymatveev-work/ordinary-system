package main

import (
	"encoding/json"
	"log"
	"net/http"
	"ordinary-system/blog/data"
	"ordinary-system/blog/model"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.
		Methods("GET").
		Path("/api/blogs").
		HandlerFunc(getAll)

	router.
		Methods("POST").
		Path("/api/blogs").
		HandlerFunc(create)

	log.Fatal(http.ListenAndServe(":8081", router))
}

func getAll(w http.ResponseWriter, r *http.Request) {
	initResponse(w)

	// TODO: obtain user ID

	articles, err := data.GetArticles(0)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(articles)
}

func create(w http.ResponseWriter, r *http.Request) {
	initResponse(w)

	// TODO: obtain user ID and an article

	article, err := data.CreateArticle(0, model.Article{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(article)
}

func initResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
