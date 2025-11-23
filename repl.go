package main

import (
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