package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func QueryHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Got parameter id:%s!\n", queryParams["id"])
	fmt.Fprintf(w, "Got parameter category:%s!", queryParams["category"])
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/articles/", QueryHandler)
	r.Queries("id", "category")
	log.Fatalln(http.ListenAndServe(":8080", r))
}
