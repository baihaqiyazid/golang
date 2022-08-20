package gojson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestDecodeFile(t *testing.T) {
	
	file, _ := os.Open("customer.json")

	decoder := json.NewDecoder(file)

	customer := &Customer{}
	decoder.Decode(customer)

	fmt.Println(customer)
}

func TestEncodeFile(t *testing.T) {

	file, _ := os.Create("customerNew.json")

	encoder := json.NewEncoder(file)

	customer := &Customer{
		FirstName: "Syam",
		LastName:  "Ling",
		Age:       21,
		Hobbies:   []string{"Gaming", "Reading", "Coding"},
		Address: []Address{
			{
				Street:  "Jl. Pengadegan",
				City:    "Jakarta Selatan",
				Country: "Indonesia",
			},
			{
				Street:  "Jl. Potlot",
				City:    "Jakarta Barat",
				Country: "Indonesia",
			},
		},
	}

	encoder.Encode(customer)
}