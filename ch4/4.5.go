package ch4

// func eliminateAdjacentDuplicates(a []string) (cleaned []string) {

// 	for i := len(a) - 1; i > 0; i-- {
// 		for a[i] == a[i-1] {
// 			a = removeSingleString(a, i-1)
// 		}
// 	}
// 	return a
// }
// func withAdjacentDuplicatesRemoved(a []string) []string {
// 	switch len(a) {
// 	case 0, 1:
// 		return a
// 	default:
// 		var filtered []string
// 		for i := range a[len(a)-1] {
// 			if a[i] != a[i+1] {
// 				filtered = append(filtered, a[i])
// 			}
// 		}
// 		return filtered

// 	}
// }

// func removeSingleString(s []string, i int) []string {
// 	for j := i; j < len(s)-1; j++ {
// 		s[j] = s[j+1]
// 	}
// 	return s[:len(s)-1]

// }
