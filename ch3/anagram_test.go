package anagram

import "testing"

func TestAnagram(t *testing.T) {
	type test = struct {
		a, b string
		want bool
	}
	for _, tt := range []test{
		{"", "", true},
		{"a", "a", true},
		{"a", "b", false},
		{"dog", "god", true},
		{"dog ", "god", false},
	} {
		if got := Anagram(tt.a, tt.b); got != tt.want {
			t.Fatalf("Anagram(%s, %s) should be %v, but is: %v", tt.a, tt.b, tt.want, got)
		}
	}
}
