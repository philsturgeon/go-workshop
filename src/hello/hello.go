package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"io"
	"log"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
  router := mux.NewRouter()
  router.HandleFunc("/", HelloServer)
  // r.HandleFunc("/products", ProductsHandler)
  // r.HandleFunc("/articles", ArticlesHandler)
	
	log.Printf("Running server on 0.0.0.0:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
