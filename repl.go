package main

import (
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	var lowerWords []string
	trimmedText := strings.Trim(text, " ")
	words := strings.Split(trimmedText, " ")
	
	for _, word := range words {
		lowerWords = append(lowerWords, strings.ToLower(word))
	}

	return lowerWords
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func help() error{
	fmt.Print("\nWelcome to the Pokedex!")
	fmt.Println("Usage: ")

	return nil
}
