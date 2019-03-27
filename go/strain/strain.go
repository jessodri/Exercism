package strain

// Ints type
type Ints []int

// Lists type
type Lists [][]int

// Strings type
type Strings []string

// Keep method for Ints
func (int Ints) Keep(f func(int) bool) Ints {

	return
}

// Discard method for Ints
func (Ints) Discard(func(int) bool) {

}

// Keep method for Lists
func (Lists) Keep(func([]int) bool) {

}

// Keep Method for Strings
func (Strings) Keep(func(string) bool) {

}
