package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type post struct {
	Title string `json:"title"`
	Tag   string `json:"tag"`
	URL   string `json:"url"`
}

func main() {
	r := mux.NewRouter()

	r.Handle("/public", public)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("ListenAndServe", nil)
	}
}

var public = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	post := &post{
		Title: "Go言語入門",
		Tag:   "Go",
		URL:   "https://rightcode.co.jp/blog/information-technology/golang-introduction-environment-1",
	}
	json.NewEncoder(w).Encode(post)
})
