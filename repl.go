package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Scottishprog/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	next          *string
	previous      *string
}

func startRepl(cfg *config) {
	userInput := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		userInput.Scan()

		userText := cleanInput(userInput.Text())
		if len(userText) == 0 {
			continue
		}

		command := userText[0]
		args := []string{}
		if len(userText) > 1 {
			args = userText[1:]
		}
		functionName, ok := getCommands()[command]
		if !ok {
			fmt.Printf("Unknown command: %s\n", command)
		} else {
			err := functionName.callback(cfg, args...)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
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
	callback    func(*config, ...string) error
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
		"explore": {
			name:        "explore",
			description: "Displays a list of pokemon for the provided area",
			callback:    commandExplore,
		},
	}
}

type MapData struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
