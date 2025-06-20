package main

import (
	"fmt"
	"os"
)

func commandExit(Config config) (config, error) {
	fmt.Printf("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return Config, nil
}
