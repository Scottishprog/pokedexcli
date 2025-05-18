package main

import "fmt"

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage: \n\n")
	for commandName, command := range getCommands() {
		fmt.Printf("%s: %s\n", commandName, command.description)
	}
	return nil
}
