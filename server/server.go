package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	petname "github.com/dustinkirkland/golang-petname"
)

func Server() {
	log.Println("Starting server")

	http.HandleFunc("/", home)
	http.HandleFunc("/petnames", petnames)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("Error ListenAndServe", err)
	}
}

func home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Welcome")
}

func petnames(w http.ResponseWriter, req *http.Request) {
	words := 2
	separator := "-"
	wordsInt, err := strconv.Atoi(req.URL.Query().Get("words"))
	if err == nil {
		words = wordsInt
	}
	fmt.Fprintf(w, petname.Generate(words, separator))

}
