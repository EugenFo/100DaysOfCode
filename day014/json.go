package main

import (
	"encoding/json"
	"fmt"
)

// it is important that the values in the struct are capitalized, if not the json encoding would not work properly
// backtickes == per-field options for the encoder
type message struct {
	Name string `json:"name"`
	Body string `json:"body"`
	Time int    `json:"time,string"`
}

type short struct {
	Origin string
	New    string
}

type result struct {
	Position []int
}

func main() {
	m := message{"Alice", "Hi", 1234552}
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

	// not really an JSON format, encoding works, but you can't exactly point to an key like m2.origin; better use struct
	/* m2 := []string{"origin:","google.de","new:","shortify"}
	fmt.Println(m2)
	n,_ := json.Marshal(m2)
	fmt.Println(string(n)) */

	// if the key from the json is not the same like the field from the struct, then this particular key will not be inserted
	text := "[{\"origin\": \"google.de\", \"new\": \"shortify\"}, {\"origin\": \"facebook.com\", \"new\": \"tinyurl\"}]"

	// normal string output
	fmt.Println(text)
	// conversion string into byte
	// bytes := []byte(text)
	fmt.Println([]byte(text))

	byt := []byte(text)

	// creates new struct and unmarshal it as a string into the struct
	var shortify []short
	json.Unmarshal(byt, &shortify)

	for i := range shortify {
		fmt.Printf("Origin = %v, New = %v\n", shortify[i].Origin, shortify[i].New)
	}
	
	
	// representation of an int array in JSON
	
	text2 := "[{\"Position\": [123, 345, 432, 21, -50]}, {\"Position\": [1, 2, 3, 4, 5, 6, 7, 8]}, {\"Position\": [54, 3, 2]}]"
	
	byt2 := []byte(text2)

	var res []result
	json.Unmarshal(byt2, &res)

	for i:= range res {
		fmt.Printf("Values of Position in Index: %v are : %v\n", i, res[i].Position)
		fmt.Printf("Length of the Array: %v\n", len(res[i].Position))
	}
}
