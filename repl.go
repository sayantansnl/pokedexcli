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

func commandExit(config *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *config, args ...string) error {
	fmt.Print("\nWelcome to the Pokedex!")
	fmt.Println("Usage: ")

	return nil
}

func commandMap(config *config, args ...string) error {
	baseUrl := "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	if config.next == "" {
		config.next = baseUrl
	}

	locationStruct, err := config.client.FetchLocationAreas(config.next)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	locations := locationStruct.Results

	for _, location := range locations {
		fmt.Printf("\n%v", location.Name)
	}
	
	config.previous = locationStruct.Previous
	config.next = locationStruct.Next

	return nil
}

func commandMapB(config *config, args ...string) error {
	if config.previous == nil {
		fmt.Print("\nyou're on the first page.")
		return nil
	}

	locationStruct, err := config.client.FetchLocationAreas(*config.previous)
	if err != nil {
		return fmt.Errorf("error fetching the previous locations, %w", err)
	}

	locations := locationStruct.Results

	for _, location := range locations {
		fmt.Printf("\n%v", location.Name)
	}

	config.next = locationStruct.Next
	config.previous = locationStruct.Previous

	return nil
}

func commandExplore(config *config, location ...string) error {
	//baseUrl := "https://pokeapi.co/api/v2/location-area/"

	if len(location) < 1 {
        return fmt.Errorf("you must provide a location area name")
    }

	locationDetailStruct, err := config.client.FetchLocationAreaDetails(location[0])
	if err != nil {
		return fmt.Errorf("error fetching pokemons: %v", err)
	}

	encounters := locationDetailStruct.PokemonEncounters

	fmt.Printf("\nExploring %s...", location)
	fmt.Print("\nFound Pokemon:")

	for _, encounter := range encounters {
		fmt.Printf("\n%v", encounter.Pokemon.Name)
	}

	return nil
}
