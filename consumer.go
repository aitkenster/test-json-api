package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type Payload struct {
	Animals Data
}

type Data struct {
	Dogs Dogs
	Horses Horses
}

type Dogs map[string]int
type Horses map[string]int

func main() {
	res, err := http.Get("http://localhost:8080")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var p Payload
	err = json.Unmarshal(body, &p)
	if err != nil {
		panic(err)
	}
	fmt.Println(p.Animals.Dogs, "\n", p.Animals.Horses)
}
