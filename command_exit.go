package main

import (
	"fmt"
	"os"
)

func commandExit(Config *config, perameter *string) error {
	fmt.Printf("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}
