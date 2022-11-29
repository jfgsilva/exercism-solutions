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

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	ch := make(chan FreqMap, len(l))
	consolidatedMap := make(FreqMap)
	ccFreq := func(s string, c chan FreqMap) {
		c <- Frequency(s)
	}
	for _, text := range l {
		go ccFreq(text, ch)
	}
	for range l {
		for key, val := range <-ch {
			consolidatedMap[key] += val
		}
	}
	return consolidatedMap
}
