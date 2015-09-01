package main

// import (
// 	"encoding/json"
// 	"io"
// 	"net/http"
// 	"log"
// )
// 
// func renderData(w http.ResponseWriter, statusCode int, rootKey string, body interface{}) {
// 	data := make(map[string]interface{}, 1)
// 	data[rootKey] = body
// 	jsonString, _ := json.Marshal(data)
// 
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(statusCode)
// 	io.WriteString(w, string(jsonString))
// }
// 
// func renderError(w http.ResponseWriter, statusCode int, errMessage error) {
// 	log.Println(errMessage)
// 
// 	data := make(map[string]interface{}, 1)
// 	data["error"] = errMessage
// 	jsonString, _ := json.Marshal(data)
// 
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(statusCode)
// 	io.WriteString(w, string(jsonString))
// }
