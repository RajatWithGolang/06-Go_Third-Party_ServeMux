package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func getName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello , %s!\n", vars["name"])
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello/{name}", getName)
	http.ListenAndServe(":8080", router)
}

//curl http://localhost:8080/hello/Gorilla