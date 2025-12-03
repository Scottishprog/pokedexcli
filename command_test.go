package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/Scottishprog/pokedexcli/internal/pokeapi"
)

func TestLocationArea_explore(t *testing.T) {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	Config := &config{
		pokeapiClient: pokeClient,
	}

	t.Run("Location Explore", func(t *testing.T) {
		_, err := Config.pokeapiClient.LocationExplore("pastoria-city-area")
		if err != nil {
			t.Errorf("cfg.pokeapiClient.LocationExplore returned error: %v", err)
		}
	})
}

func TestExplore(t *testing.T) {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	Config := &config{
		pokeapiClient: pokeClient,
	}

	err := commandExplore(Config, "pastoria-city-area")
	if err != nil {
		fmt.Println("Error - commandExplore:", err)
	}
}
