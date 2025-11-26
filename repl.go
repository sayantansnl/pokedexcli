package main

import (
	"fmt"
	"math/rand/v2"
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

	if len(location) < 1 {
        return fmt.Errorf("you must provide a location area name")
    }

	locationDetailStruct, err := config.client.FetchLocationAreaDetails(strings.ToLower(location[0]))
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

func commandCatch (config *config, pokemonName ...string) error {
	if len(pokemonName) < 1 {
		return fmt.Errorf("you must provide a pokemon name")
	}

	pokemonDetailStruct, err := config.client.FetchPokemon(strings.ToLower(pokemonName[0]))
	if err != nil {
		return fmt.Errorf("error fetching pokemon details: %w", err)
	}

	name := pokemonDetailStruct.Name
	baseExperience := pokemonDetailStruct.BaseExperience

	if _, ok := config.pokedex[name]; ok {
		return fmt.Errorf("%s is already caught", name)
	}

	fmt.Printf("\nThrowing a Pokeball at %s...", name)

	ballPower := rand.IntN(2 * baseExperience)

	if ballPower < baseExperience {
		fmt.Printf("\n%s escaped!", name)
	} else {
		fmt.Printf("\n%s was caught!", name)
		config.pokedex[name] = pokemonDetailStruct
	}

	return nil
}

func commandInspect (config *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("please enter the name of the pokemon you want to inspect")
	}
	
	if _, ok := config.pokedex[args[0]]; !ok {
		return fmt.Errorf("you haven't caught that pokemon")
	}

	caughtPokemon := config.pokedex[args[0]]

	name := caughtPokemon.Name
	height := caughtPokemon.Height
	weight := caughtPokemon.Weight
	stats := caughtPokemon.Stats
	types := caughtPokemon.Types

	fmt.Printf("\nName: %s", name)
	fmt.Printf("\nHeight: %d", height)
	fmt.Printf("\nWeight: %d", weight)
	fmt.Printf("\nStats:")

	for _, stat := range stats {
		fmt.Printf("\n-%v: %v", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Printf("\nTypes:")

	for _, t := range types {
		fmt.Printf("\n-%v", t.Type.Name)
	}
	
	return nil

}

func commandPokedex (config *config, args ...string) error {
	if len(config.pokedex) == 0 {
		return fmt.Errorf("there are no caught pokemons in your pokeballs")
	}

	for key, _ := range config.pokedex {
		fmt.Printf("\n-%s", key)
	}
	return nil
}
