package main

import (
	"github.com/gorilla/mux"
	j "github.com/ricardolonga/jsongo"
	"net/http"
	"io"
	"log"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	json := j.Object().Put("name", "Ricardo Longa").
		Put("idade", 28).
		Put("owner", true).
		Put("skills", j.Array().Put("Golang").
		                       Put("Android"))

	io.WriteString(w, json.Indent())
}

func main() {
  router := mux.NewRouter()
  router.HandleFunc("/", HelloServer)
  // r.HandleFunc("/products", ProductsHandler)
  // r.HandleFunc("/articles", ArticlesHandler)
	
	log.Printf("Running server on 0.0.0.0:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
