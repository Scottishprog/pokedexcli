package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	userInput := bufio.NewScanner(os.Stdin)
	var Conf config
	for {
		fmt.Print("Pokedex > ")
		userInput.Scan()

		userText := cleanInput(userInput.Text())
		if len(userText) == 0 {
			continue
		}

		command := userText[0]
		functionName, ok := getCommands()[command]
		if !ok {
			fmt.Printf("Unknown command: %s\n", command)
		} else {
			err := functionName.callback(&Conf)
			if err != nil {
				return
			}
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},

		"map": {
			name:        "map",
			description: "Displays the next 20 areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 areas",
			callback:    commandMapb,
		},
	}
}

type config struct {
	next     string
	previous string
}
