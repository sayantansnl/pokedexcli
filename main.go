package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/sayantansnl/pokedexcli/internal/pokeapi"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	cfg := &config{
		client: pokeapi.NewClient("https://pokeapi.co/api/v2", 5*time.Second),
	}

	for {
		fmt.Print("\nPokedex > ")
		if scanner.Scan() {
			input := scanner.Text()
			cleanedInput := cleanInput(input)

			if len(cleanedInput) == 0 {
				continue
			}

			command := cleanedInput[0]
			
			value, ok := commandRegistry[command]
			if !ok {
				fmt.Println("Unknown command")
				continue
			}

			if err := value.callback(cfg); err != nil {
				fmt.Println(err)
			}

			if value.name == "help" {
				for key, value := range commandRegistry {
					fmt.Printf("\n%v: %v", key, value.description)
				}
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("\nError in printing: %v", err)
		}
	}
}