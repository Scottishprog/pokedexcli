package main

import (
	"github.com/Scottishprog/pokedexcli/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	Config := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(Config)
}
