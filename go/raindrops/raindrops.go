package raindrops

import "fmt"

// Convert function returns a specific string based on a numbers factors
func Convert(num int) (word string) {

	pling := "Pling"
	plang := "Plang"
	plong := "Plong"

	factor3 := num%3 == 0
	factor5 := num%5 == 0
	factor7 := num%7 == 0

	result := ""

	if factor3 {
		result += pling
	}
	if factor5 {
		result += plang
	}
	if factor7 {
		result += plong
	}
	if !factor3 && !factor5 && !factor7 {
		result = fmt.Sprintf("%v", num)
	}

	return result
}
