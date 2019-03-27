package hamming

import "errors"

// Distance function to calculate the Hamming difference between two DNA strands
func Distance(a, b string) (int, error) {

	distance := 0

	if len(a) != len(b) {
		return 0, errors.New("String input a and string inpout b must be the same length")
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
