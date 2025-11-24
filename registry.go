package main

type cliCommand struct {
	name          string
	description   string
	callback      func(*config) error
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
}