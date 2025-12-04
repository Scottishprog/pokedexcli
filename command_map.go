package main

import (
	"fmt"
)

func commandMap(cfg *config, parameter *string) error {
	shallowmapdata, err := cfg.pokeapiClient.LocationList(cfg.next)
	if err != nil {
		return err
	}

	cfg.next = shallowmapdata.Next
	cfg.previous = shallowmapdata.Previous

	for _, result := range shallowmapdata.Results {
		fmt.Println(result.Name)
	}

	return nil

}

func commandMapb(cfg *config, perameter *string) error {
	if cfg.previous == nil {
		fmt.Println("On the first page")
		return nil
	}

	shallowmapdata, err := cfg.pokeapiClient.LocationList(cfg.previous)
	if err != nil {
		return err
	}

	cfg.next = shallowmapdata.Next
	cfg.previous = shallowmapdata.Previous

	for _, result := range shallowmapdata.Results {
		fmt.Println(result.Name)
	}

	return nil

}
