package main

import (
	"net/http"
	"fmt"
	"encoding/json"
)

type Payload struct {
	Animals  Data
}

type Data struct {
	Dogs Dogs
	Horses Horses
}

type Dogs map[string]int
type Horses map[string]int

func serveResponse(w http.ResponseWriter, r *http.Request) {
	response, err := getJsonResponse()
	if err != nil {
		panic(err)
	}

	fmt.Fprint(w, string(response))
}

func main() {
	http.HandleFunc("/", serveResponse)
	http.ListenAndServe("localhost:8080", nil)
}

func getJsonResponse() ([]byte, error) {
	dogs := make(map[string]int)
	dogs["Poodles"] = 17
	dogs["Spaniels"] = 50
	dogs["Dobermans"] = 4

	horses := make(map[string]int)
	horses["Thoroghbreds"] = 2
	horses["Shetlands"] = 77
	horses["Connemaras"] = 23

	d := Data{dogs, horses}
	p := Payload{d}

	return json.MarshalIndent(p, "", "  ")
}
