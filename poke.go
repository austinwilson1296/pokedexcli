package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Struct to handle API response
type pokeAPI struct {
	Count    int      `json:"count"`
	Next     string   `json:"next"`
	Previous *string  `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// Makes a GET request to the pokeAPI and returns a page of locations
func getMap(url string) (pokeAPI, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return pokeAPI{}, fmt.Errorf("error making request: %w", err)
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return pokeAPI{}, err
	}
	defer res.Body.Close()

	// Decode JSON response into pokeAPI struct
	var poke pokeAPI
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&poke)
	if err != nil {
		return pokeAPI{}, err
	}

	return poke, nil
}
