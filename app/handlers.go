package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Customer struct {
	Name    string   `json:"name"`
	City    string   `json:"city"`
	ZipCode string   `json:"zipCode"`
	Posts   []string `json:"posts"`
}

func fun(w http.ResponseWriter, r *http.Request) {

	customer := []Customer{
		{Name: "Saif", City: "Baghdad", ZipCode: "00964", Posts: []string{
			"Post1sdsd",
			"Post2",
			"Post3",
			"Post4",
			"Post5",
		}},
		{Name: "Saif1", City: "Baghdad1", ZipCode: "009641", Posts: []string{
			"Post1",
			"Post2",
			"Post3",
			"Post4",
			"Post5",
		}},
		{Name: "Saif2", City: "Baghdad2", ZipCode: "009642", Posts: []string{
			"Post1",
			"Post2",
			"Post3",
			"Post4",
			"Post5",
		}},
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	fmt.Fprintf(w, vars["id"])
}
