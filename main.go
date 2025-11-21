package main

import (
	"time"

	"github.com/Scottishprog/pokedexcli/internal/pokeapi"
	"github.com/Scottishprog/pokedexcli/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	Config := &config{
		pokeapiClient: pokeClient,
	}

	cache := pokecache.NewCache(5 * time.Second)
	startRepl(Config)
}
