package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

//Article adalah data
type Article struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
}

var post []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello From Golang")
}

func getAllArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func getSingleArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range post {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	http.NotFound(w, r)
}

func addArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var posts Article

	json.NewDecoder(r.Body).Decode(&posts)
	posts.ID = strconv.Itoa(rand.Intn(100))
	post = append(post, posts)
	json.NewEncoder(w).Encode(posts)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, v := range post {
		if v.ID == params["id"] {
			post = append(post[:i], post[i+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(post)
}

func editArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var posts Article

	for i, v := range post {
		if v.ID == params["id"] {
			post = append(post[:i], post[i+1:]...)
			json.NewDecoder(r.Body).Decode(&posts)
			posts.ID = params["id"]
			post = append(post, posts)
			return
		}
	}

	json.NewEncoder(w).Encode(post)

}

func main() {

	// add data to slice[]artice or fakedb

	post = append(post, Article{
		ID:          "1",
		Title:       "How To Write Your First Program in Go",
		Description: "The “Hello, World!” program is a classic and time-honored tradition in computer programming. It's a simple and complete first program for beginners, and it's a good way to make sure your environment is properly configured. This tutorial will walk you through creating this program in Go. ",
		Author:      "John Snow",
	})

	post = append(post, Article{
		ID:          "2",
		Title:       "Understanding the GOPATH",
		Description: "This article will walk you through understanding what the `GOPATH` is, how it works, and how to set it up. This is a crucial step for setting up a Go development environment, as well as understanding how Go finds, installs, and builds source files. ",
		Author:      "Don Jon",
	})

	// Route handles
	r := mux.NewRouter()
	r.HandleFunc("/", homePage)
	r.HandleFunc("/post", getAllArticle).Methods("GET")
	r.HandleFunc("/post/{id}", getSingleArticle).Methods("GET")
	r.HandleFunc("/post", addArticle).Methods("POST")
	r.HandleFunc("/post/{id}", editArticle).Methods("PUT")
	r.HandleFunc("/post/{id}", deleteArticle).Methods("DELETE")

	//Start the server
	log.Fatal(http.ListenAndServe(":8000", r))

}
