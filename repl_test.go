package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input     string
		expected  []string
	}{
		{
			input: "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "Charmander Bulbasaur Pikachu",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input: "Goku Vegeta Gohan",
			expected: []string{"goku", "vegeta", "gohan"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length of actual slice: %v, Length of expected slice: %v", len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Expected word: %s, Actual word: %s", expectedWord, word)
			}
		}
	}
}