package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\nPokedex > ")
		if scanner.Scan() {
			input := scanner.Text()
			cleanedInput := cleanInput(input)
			fmt.Printf("Your command was: %v", cleanedInput[0])
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("\nError in printing: %v", err)
		}
	}
}