package main

import (
	"errors"
	"fmt"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("Usage: catch <pokemon>")
	}

	pokemon := args[0]
	pokemondata, err := cfg.pokeapiClient.PokemonData(pokemon)
	if err != nil {
		return err
	}

	fmt.Println(pokemondata)

	return nil
}
