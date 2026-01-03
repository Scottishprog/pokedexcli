package main

import (
	"time"

	"github.com/Scottishprog/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	Config := &config{
		pokeapiClient: pokeClient,
		Pokemon:       map[string]pokeapi.PokemonData{},
	}

	startRepl(Config)
}
