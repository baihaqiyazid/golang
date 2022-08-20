package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
}

type Customer struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Age       int       `json:"age"`
	Hobbies   []string  `json:"hobbies"`
	Address   []Address `json:"address"`
}

func TestJSONArray(t *testing.T) {

	customer := Customer{
		FirstName: "Yazid",
		LastName:  "Baihaqi",
		Age:       21,
		Hobbies:   []string{"Gaming", "Reading", "Coding"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {

	jsonString := `{"first_name":"Yazid","last_name":"Baihaqi","age":21,"hobbies":["Gaming","Reading","Coding"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T) {

	customer := Customer{
		FirstName: "Yazid",
		LastName:  "Baihaqi",
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

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {

	jsonString := `{"first_name":"Yazid",
					"last_name":"Baihaqi",
					"age":21,
					"hobbies":["Gaming","Reading","Coding"],
					"address":
						[
							{"street":"Jl. Pengadegan","city":"Jakarta Selatan","country":"Indonesia"},
							{"street":"Jl. Potlot","city":"Jakarta Barat","country":"Indonesia"}
						]
					}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println("FirstName:", customer.FirstName)
	fmt.Println("LastName :", customer.LastName)
	fmt.Println("Age      :", customer.Age)
	fmt.Println("Hobibies :", customer.Hobbies)
	for i, v := range customer.Address {
		fmt.Printf("Alamat %d :", i)
		fmt.Println(v)
	}

	// fmt.Println("Alamat 2 :",customer.Address[1])
}
