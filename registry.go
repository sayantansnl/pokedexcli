package main

type cliCommand struct {
	name          string
	description   string
	callback      func(*config, ...string) error
}

var commandRegistry = map[string]cliCommand{
	"exit": {
		name: "exit",
		description: "Exit the Pokedex",
		callback: commandExit,
	},
	"help": {
		name: "help",
		description: "Displays a help message",
		callback: commandHelp,
	},
	"map": {
		name: "map",
		description: "Displays the names of 20 locations",
		callback: commandMap,
	},
	"mapb": {
		name: "mapb",
		description: "Displays the names of the 20 previous locations",
		callback: commandMapB,
	},
	"explore": {
		name: "explore",
		description: "Displays the names of Pokemons available in a location",
		callback: commandExplore,
	},
	"catch" : {
		name: "catch",
		description: "Attempts to catch the Pokemon given by a name",
		callback: commandCatch,
	},
	"inspect": {
		name: "inspect",
		description: "Inspects the details of the caught pokemon",
		callback: commandInspect,
	},
	"pokedex": {
		name: "pokedex",
		description: "Displays the names of the caught pokemons",
		callback: commandPokedex,
	},
}