// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) (completeRhyme []string) {
	if len(rhyme) != 0 {
		for i := 0; i < len(rhyme)-1; i++ {
			completeRhyme = append(completeRhyme, fmt.Sprintf("For want of a %v the %v was lost.", rhyme[i], rhyme[i+1]))
		}
		completeRhyme = append(completeRhyme, fmt.Sprintf("And all for the want of a %v.", rhyme[0]))
	}
	return
}
