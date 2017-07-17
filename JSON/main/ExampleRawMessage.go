package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func ExampleRawMessage_marshal() {
	h := json.RawMessage(`{"precomputed": true}`)

	c := struct {
		Header *json.RawMessage `json:"header"`
		Body   string           `json:"body"`
	}{Header: &h, Body: "Hello Gophers!"}

	b, err := json.MarshalIndent(&c, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)

	// Output:
	// {
	// 	"header": {
	// 		"precomputed": true
	// 	},
	// 	"body": "Hello Gophers!"
	// }
}

func main()  {
	ExampleRawMessage_marshal()
}
