package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"ordinary-system/blog/data"
	"ordinary-system/blog/model"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.
		Methods("GET").
		Path("/api/users/{userID}/articles").
		HandlerFunc(getAll)

	router.
		Methods("POST").
		Path("/api/users/{userID}/articles").
		HandlerFunc(create)

	log.Fatal(http.ListenAndServe(":8081", router))
}

func getAll(w http.ResponseWriter, r *http.Request) {
	initResponse(w)

	vars := mux.Vars(r)
	userID, err := strconv.ParseInt(vars["userID"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	articles, err := data.GetArticles(userID)

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

	vars := mux.Vars(r)
	userID, err := strconv.ParseInt(vars["userID"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	article, err := getArticleFromBody(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	article, err = data.CreateArticle(userID, article)
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

func getArticleFromBody(body io.Reader) (model.Article, error) {
	rawBody, _ := ioutil.ReadAll(body)
	var a model.Article
	if err := json.Unmarshal(rawBody, &a); err != nil {
		return model.Article{}, err
	}
	return a, nil
}
