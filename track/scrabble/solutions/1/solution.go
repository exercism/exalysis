package scrabble

import "strings"

func scoreLetter(char rune) int {
	switch char {
	case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
		return 1
	case 'D', 'G':
		return 2
	case 'B', 'C', 'M', 'P':
		return 3
	case 'F', 'H', 'V', 'W', 'Y':
		return 4
	case 'K':
		return 5
	case 'J', 'X':
		return 8
	case 'Q', 'Z':
		return 10
	default:
		return 0
	}
}

// Score calculates scrabble score for the given word
func Score(word string) int {
	word = strings.ToUpper(word)
	sum := 0
	for _, char := range word {
		sum += scoreLetter(char)
	}
	return sum
}
