package main

import (
	"fmt"
	"log"
	"net/http"
)
func main(){
	http.HandleFunc("/", worker)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
func worker(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "successifull build a simple server")
}