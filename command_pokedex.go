package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {

	if len(cfg.Pokemon) == 0 {
		fmt.Println("No pokemon")
	}

	for _, pokemon := range cfg.Pokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
