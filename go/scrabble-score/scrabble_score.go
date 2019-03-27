// Package scrabble returns the scrabble score for each word
package scrabble

import "unicode"

// processLetter assigns a value to each letter
func processLetter(letter rune) int {
	switch letter {
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

// Score adds up the total scrabble score for a word
func Score(word string) int {
	score := 0

	for _, r := range word {
		score += processLetter(unicode.ToUpper(r))
	}
	return score
}
