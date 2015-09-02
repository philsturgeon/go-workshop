package main

import (
	"log"
	"errors"
	"github.com/gorilla/mux"
	"gopkg.in/olivere/elastic.v2"
	"net/http"
	"time"
	"os"
)

var elasticClient *elastic.Client

func init() {
	if err := establishElasticConnection(); err != nil {
		log.Panicln(err)
	}
}

func vehicleSearch(w http.ResponseWriter, req *http.Request) {
	vehicles, _ := searchForVehicles(elasticClient, req.FormValue("search"))
	
	renderData(w, http.StatusOK, "vehicles", vehicles)
}

func main() {
	router := mux.NewRouter()

	vehicleSearchHandler := http.HandlerFunc(vehicleSearch)
	router.Handle("/vehicles", corsHandler(vehicleSearchHandler))
	
	port := os.Getenv("PORT")
	
	log.Printf("Running server on 0.0.0.0:" + port)
	log.Fatal(http.ListenAndServe(":" + port, router))
}

func createElasticClient() (*elastic.Client, error) {
	client, err := elastic.NewClient()
	if err != nil {
		return nil, err
	}
	if res, err := client.ClusterHealth().WaitForStatus("red").Timeout("15s").Do(); err != nil {
		return nil, err
	} else if res.TimedOut {
		return nil, errors.New("time out waiting for cluster status red")
	}
	return client, nil
}

func establishElasticConnection() error {
	var err error
	for i := 0; i < 10; i++ {
		elasticClient, err = createElasticClient()
		if err != nil {
			log.Println("Sleeping one second to wait for elasticsearch")
			time.Sleep(time.Second)
		} else {
			return nil
		}
	}
	return err
}


func corsHandler(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    next.ServeHTTP(w, r)
  })
}
