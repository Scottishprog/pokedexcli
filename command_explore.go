package main

import (
	"fmt"
)

func commandExplore(cfg *config, location string) error {
	deepmapdata, err := cfg.pokeapiClient.LocationExplore(location)
	if err != nil {
		return err
	}

	for _, pokemon := range deepmapdata.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}
