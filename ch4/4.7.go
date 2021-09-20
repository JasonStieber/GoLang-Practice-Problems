package ch4

func reverseCharactersInPlace(a []byte) []byte {
	// 4 ifs based on byte length
	s := ""
	for i := 0; i < len(a); i++ {
		if a[i] < 128 { // this is an ASCII character
			s = string(a[i]) + s
		} else if a[i] >= 192 && a[i] <= 223 { // two byte rune
			s = string(a[i:i+2]) + s
			i++
		} else if a[i] >= 224 && a[i] <= 239 { // three byte rune
			s = string(a[i:i+3]) + s
			i += 2
		} else if a[i] >= 240 && a[i] <= 247 { // four byte rune
			s = string((a[i : i+4])) + s
			i += 3
		}
	}
	return []byte(s)

}
