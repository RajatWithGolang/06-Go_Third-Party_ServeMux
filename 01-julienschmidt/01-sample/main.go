package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)
func getName(w http.ResponseWriter, r *http.Request, param httprouter.Params){
fmt.Fprintf(w, "Hello , %s!\n", param.ByName("name"))
}
func postName(w http.ResponseWriter, r *http.Request, param httprouter.Params){
fmt.Fprintf(w, "Hello , %s!\n", param.ByName("name"))
}
func main(){
	router := httprouter.New()
	router.GET("/hello/:name",getName)
	router.POST("/hello/:name",postName)
	http.ListenAndServe(":8080",router)
}