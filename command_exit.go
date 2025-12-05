package main

import (
	"fmt"
	"os"
)

func commandExit(Config *config, args ...string) error {
	fmt.Printf("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}
