package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}

	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency is a function that runs Frequency() and send results to different channels
func ConcurrentFrequency(s []string) FreqMap {
	c := make(chan FreqMap)
	for _, r := range s {
		go func(s string) {
			c <- Frequency(s)
		}(r)
	}

	uberMap := FreqMap{}

	for range s {

		res := <-c
		for k, v := range res {
			uberMap[k] += v
		}
	}
	return uberMap
}
