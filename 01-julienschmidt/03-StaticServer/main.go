package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
)

func main(){
	router := httprouter.New()
	router.ServeFiles("/Public/*filepath",http.Dir("."))
	log.Fatal(http.ListenAndServe(":8080", router))

}