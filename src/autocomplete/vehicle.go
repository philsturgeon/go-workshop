package main

// import (
// 	"gopkg.in/olivere/elastic.v2"
// 	"encoding/json"
// 	"net/http"
// 	"errors"
// 	"log"
// )
// 
// const vehicleIndex = "vehicle"
// 
// type Vehicle struct {
// 	Make           string `json:"make"`
// 	Year      		 int    `json:"year"`
// 	Model          string `json:"model"`
// 	PrimaryFuel    string `json:"primary_fuel"`
// 	VehicleClass   string `json:"vehicle_class"`
// 	CityMPG        string `json:"city_mpg"`
// 	HighwayMPG     string `json:"highway_mpg"`
// 	CombMPG        string `json:"comb_mpg"`
// }
// 
// func (vehicle Vehicle) GetJson() (string, error) {
// 	out, err := json.Marshal(vehicle)
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(out), nil
// }
// 
// func searchForVehicles(client *elastic.Client, searchTerm string) (vehicles []Vehicle, err error) {
// 	matchQuery := elastic.NewMatchQuery("_all", searchTerm).Operator("and")
// 
// 	searchResult, err := client.Search().
// 		Index(vehicleIndex).
// 		Query(&matchQuery).
// 		Sort("_score", false).
// 		Sort("year", false).
// 		// Debug(true).Pretty(true).
// 		Do()
// 	if err != nil {
// 		// Handle error
// 		return nil, err
// 	}
// 
// 	// Number of hits
// 	if searchResult.Hits == nil {
// 		vehicles = make([]Vehicle, 0)
// 		return
// 	}
// 
// 	if vehicleCount := len(searchResult.Hits.Hits); vehicleCount > 0 {
// 		vehicles = make([]Vehicle, vehicleCount)
// 	} else {
// 		vehicles = make([]Vehicle, 0)
// 		return
// 	}
// 
// 	for i, hit := range searchResult.Hits.Hits {
// 		var resultVehicle Vehicle
// 		err := json.Unmarshal(*hit.Source, &resultVehicle)
// 		if err != nil {
// 			log.Printf("Failed to deserialize: %+v", resultVehicle)
// 			continue
// 		}
// 		vehicles[i] = resultVehicle
// 	}
// 	return
// }
// 
// func vehicleSearch(w http.ResponseWriter, req *http.Request) {
// 	vehicles, err := searchForVehicles(elasticClient, req.FormValue("search"))
// 	if err != nil {
// 		renderError(w, http.StatusInternalServerError, errors.New("Shit went wrong"))
// 		return
// 	}
// 	// if 
// 	renderData(w, http.StatusOK, "vehicles", vehicles)
// }
