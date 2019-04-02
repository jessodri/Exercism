package robotname

import "math/rand"

// Robot Struct
type Robot struct {
	name string
}

// RandomString generates random string
func RandomString() string {
	var letter = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var number = []rune("0123456789")

	b := make([]rune, 2)
	for i := 0; i < 2; i++ {
		b[i] = letter[rand.Intn(len(letter))]
	}

	c := make([]rune, 3)
	for i := 0; i < 3; i++ {
		c[i] = number[rand.Intn(len(number))]
	}

	return string(append(b, c...))
}

// Name returns the robot name
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	r.name = RandomString()
	return r.name, nil
}

// Reset resets the robot name
func (r *Robot) Reset() string {
	r.name = ""
	return r.name
}
