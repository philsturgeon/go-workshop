package main

import (
	"log"
	// "errors"
	"github.com/gorilla/mux"
	// "gopkg.in/olivere/elastic.v2"
	"net/http"
	"io"
	// "time"
)

// var elasticClient *elastic.Client

func vehicleSearch(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", vehicleSearch)

	log.Printf("Running server on 0.0.0.0:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// func createElasticClient() (*elastic.Client, error) {
// 	// Check for an elastic URL, which is probably missing in development
// 	// elasticURL := os.Getenv("ELASTIC_URL")
// 	// if elasticURL == "" {
// 	// 	elasticURL = "http://localhost:9200"
// 	// }
// 	client, err := elastic.NewClient()
// 	if err != nil {
// 		return nil, err
// 	}
// 	if res, err := client.ClusterHealth().WaitForStatus("red").Timeout("15s").Do(); err != nil {
// 		return nil, err
// 	} else if res.TimedOut {
// 		return nil, errors.New("time out waiting for cluster status red")
// 	}
// 	return client, nil
// }

// func establishElasticConnection() error {
// 	var err error
// 	for i := 0; i < 10; i++ {
// 		elasticClient, err = createElasticClient()
// 		if err != nil {
// 			log.Println("Sleeping one second to wait for elasticsearch")
// 			time.Sleep(time.Second)
// 		} else {
// 			return nil
// 		}
// 	}
// 	return err
// }

// func init() {
// elasticClient, err = createElasticClient()

// if err := establishElasticConnection(); err != nil {
// 	log.Panicln(err)
// }
// }

// func corsHandler(next http.Handler) http.Handler {
//   return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Access-Control-Allow-Origin", "*")
//     next.ServeHTTP(w, r)
//   })
// }
