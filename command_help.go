package main

import "fmt"

func commandHelp(Config *config) error {
	fmt.Println()
	fmt.Printf("Welcome to the Pokedex!\nUsage: \n\n")
	for commandName, command := range getCommands() {
		fmt.Printf("%s: %s\n", commandName, command.description)
	}

	fmt.Println("Next: ", Config.next, "\nPrevious: ", Config.previous)

	return nil
}
