package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONMap(t *testing.T) {

	product := map[string]interface{} {
		"id" : "p980",
		"name": "Hoddie Black",
		"price": 300000,
	}

	byte, _ := json.Marshal(product)
	fmt.Println(string(byte))
}

func TestJSONMapDecode(t *testing.T) {

	jsonString := `{"id":"p980","name":"Hoddie Black","price":300000}`
	jsonBytes := []byte(jsonString)

	var products map[string]interface{}

	err := json.Unmarshal(jsonBytes, &products)
	if err != nil {
		panic(err)
	}

	for i, v := range products {
		fmt.Println(i, v)
	}
}