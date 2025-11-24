package main

type cliCommand struct {
	name          string
	description   string
	callback      func() error
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
		callback: help,
	},
}