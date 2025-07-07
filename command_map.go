package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func commandMap(Config *config) error {
	// get map information from:

	// get this code working - produces a named struct that can be stepped through to print out all the map
	// locations. - DONE
	// Pull all API related code to an internal package, this func will
	// only deal with requesting the struct, and display. - WIP
	// Get the Next/Previous struct parameters working with the
	// function callback. - DONE

	var address string
	if Config.next == "" {
		address = "https://pokeapi.co/api/v2/location-area/"
	} else {
		address = Config.next
	}

	mapdata := LocationAreaList(address)
	res, err := http.Get(address)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	type mapData struct {
		Count    int    `json:"count"`
		Next     string `json:"next"`
		Previous string `json:"previous"`
		Results  []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"results"`
	}

	mapdata := mapData{}
	err = json.Unmarshal(body, &mapdata)
	if err != nil {
		log.Fatal(err)
	}
	Config.next = mapdata.Next
	Config.previous = mapdata.Previous

	fmt.Println("Next: ", Config.next, "\nPrevious: ", Config.previous)
	for _, result := range mapdata.Results {
		fmt.Println(result.Name)
	}

	//fmt.Println(mapdata)
	return nil

}
