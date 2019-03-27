package accumulate

// Accumulate performs tbe same task on each item in a collection
func Accumulate(collection []string, operation func(string) string) []string {
	result := make([]string, len(collection))
	for char, word := range collection {
		result[char] = operation(word)
	}
	return result
}
