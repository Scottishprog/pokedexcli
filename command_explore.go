package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("explore requires 1 argument: explore <command>")
	}
	location := args[0]
	deepmapdata, err := cfg.pokeapiClient.LocationExplore(location)
	if err != nil {
		return err
	}

	fmt.Println("\n")
	for _, pokemon := range deepmapdata.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}
