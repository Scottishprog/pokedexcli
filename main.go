package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	result := make([]string, 0)
	splitSlice := strings.Fields(text)
	for _, word := range splitSlice {
		result = append(result, strings.ToLower(word))
	}
	//result = append(result, strings.Split(text, " "))

	return result
}

func main() {
	fmt.Print("Hello, World! \n")
}
