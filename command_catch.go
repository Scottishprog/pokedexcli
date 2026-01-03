package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("Usage: catch <pokemon>")
	}
	if len(args) != 1 {
		return errors.New("Usage: catch <pokemon>")
	}

	pokemon := args[0]
	pokemondata, err := cfg.pokeapiClient.Cetch_Pokemmon(pokemon)
	if err != nil {
		return err
	}

	res := rand.Intn(pokemondata.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemondata.Name)
	if res > 40 {
		fmt.Printf(" %s escaped!\n", pokemondata.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon)

	cfg.Pokemon[pokemon] = pokemondata

	return nil
}
