package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Food struct
type Food struct {
	ID       string  `json:"id"`
	Group    string  `json:"group"`
	Subgroup string  `json:"subgroup"`
	Name     string  `json:"Name"`
	Origin   *Origin `json:"origin"`
}

//Origin struct
type Origin struct {
	Country string `json:"country"`
	City    string `json:"city"`
}

//food slice
var foods []Food

//get all foods
func getFoods(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(foods)
}

//get food by id
func getFood(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //getting params for food

	//loop through to find id
	for _, item := range foods {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Food{})
}

//add food
func createFood(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var food Food
	_ = json.NewDecoder(r.Body).Decode(&food)
	food.ID = strconv.Itoa(rand.Intn(10000000)) //Mock ID, not safe
	foods = append(foods, food)
	json.NewEncoder(w).Encode(food)
}

//update food by id uses delete and create
func updateFood(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //getting param to update
	//loop through to find id
	for index, item := range foods {
		if item.ID == params["id"] {
			foods = append(foods[:index], foods[index+1:]...) //this deletes.. idk how
			var food Food
			_ = json.NewDecoder(r.Body).Decode(&food)
			food.ID = params["id"]
			foods = append(foods, food)
			json.NewEncoder(w).Encode(food)
			return
		}
		json.NewEncoder(w).Encode(foods)
	}

}

//delete food by id
func deleteFood(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //getting param to delete
	//loop through to find id
	for index, item := range foods {
		if item.ID == params["id"] {
			foods = append(foods[:index], foods[index+1:]...) //this deletes.. idk how
			break
		}
		json.NewEncoder(w).Encode(foods)
	}
}

func main() {
	//init router
	r := mux.NewRouter()

	//test data @ add db in futurer
	foods = append(foods, Food{ID: "1", Group: "Dairy", Subgroup: "Cheese", Name: "Parmigiano Reggiano",
		Origin: &Origin{Country: "Italy", City: "Parma"}})
	foods = append(foods, Food{ID: "2", Group: "Dairy", Subgroup: "Cheese", Name: "Brie",
		Origin: &Origin{Country: "France", City: "Seine-et-Marne"}})
	foods = append(foods, Food{ID: "3", Group: "Dairy", Subgroup: "Cheese", Name: "Pecorino",
		Origin: &Origin{Country: "Italy", City: "Sardinia"}})

	//handle router / endpoints
	r.HandleFunc("/api/foods", getFoods).Methods("GET")
	r.HandleFunc("/api/foods/{id}", getFood).Methods("GET")
	r.HandleFunc("/api/foods", createFood).Methods("POST")
	r.HandleFunc("/api/foods/{id}", updateFood).Methods("PUT")
	r.HandleFunc("/api/foods/{id}", deleteFood).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
