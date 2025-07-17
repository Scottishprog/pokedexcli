package main

import (
	"fmt"
)

func commandMap(cfg *config) error {
	shallowmapdata, err := cfg.pokeapiClient.LocationList(cfg.next)
	if err != nil {
		return err
	}

	cfg.next = shallowmapdata.Next
	cfg.previous = shallowmapdata.Previous

	fmt.Println("Next: ", &cfg.next, "\nPrevious: ", shallowmapdata.Previous)
	for _, result := range shallowmapdata.Results {
		fmt.Println(result.Name)
	}

	return nil

}
