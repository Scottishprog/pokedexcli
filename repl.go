package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	userInput := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		userInput.Scan()

		userText := cleanInput(userInput.Text())
		if len(userText) == 0 {
			continue
		}

		commandName := userText[0]
		fmt.Printf("Your command was: %s\n", commandName)

	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
