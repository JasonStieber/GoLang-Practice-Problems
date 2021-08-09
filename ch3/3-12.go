package ch3

// 3.12: write a function that reports whether two strings are anagrams of each other,
// that is ,they contain the same letters in a different order.
func Anagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	aMap := make(map[byte]int)
	bMap := make(map[byte]int)
	for i := range a {
		aMap[a[i]]++
	}
	for i := range b {
		bMap[b[i]]++
	}
	if len(aMap) != len(bMap) {
		return false
	}
	for k := range aMap {
		if aMap[k] != bMap[k] {
			return false
		}
	}
	return true
}
