// Package scrabble implements the Exercism Scrabble Score solution.
package scrabble

import (
	"log"
	"strings"
)

/*
Letter                           Value
A, E, I, O, U, L, N, R, S, T       1
D, G                               2
B, C, M, P                         3
F, H, V, W, Y                      4
K                                  5
J, X                               8
Q, Z                               10
*/
var letterValues = []int{
	1, 3, 3, 2, 1, 4, 2, // A-G
	4, 1, 8, 5, 1, 3, 1, 1, 3, // H-P
	10, 1, 1, // Q-S
	1, 1, 4, // T-V
	4, 8, 4, 10, // W-Z
}

// Define bounds checking values.
var asciiA = 'A'
var asciiZ = 'Z'

// Score returns the Scrabble score of a provided word.
func Score(input string) (score int) {
	r := []rune(strings.ToUpper(input))

	for _, letter := range r {
		// Score can only evaluate ASCII characters.
		if letter < asciiA || letter > asciiZ {
			log.Printf("Invalid letter: %s\n", string(letter))
			return
		}

		score += letterValues[letter-asciiA]
	}

	return
}
