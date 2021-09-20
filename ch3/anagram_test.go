package ch3

import "testing"

func TestComma(t *testing.T) {
	type test = struct {
		a, want string
	}
	for _, tt := range []test{
		{"100001", "100,001"},
		{"101", "101"},
		{"2", "2"},
		{"1234", "1,234"},   // off by one
		{"12345", "12,345"}, // off by two
		{"12345678910121314", "12,345,678,910,121,314"},
	} {
		if got := Comma(tt.a); got != tt.want {
			t.Errorf("Comma(%s) should return %v: but it returned: %v", tt.a, tt.want, got)
		}
	}
}

func TestEnhancedComma(t *testing.T) {
	type test = struct {
		a, want string
	}
	for _, tt := range []test{
		{"-1", "-1"},
		{"+10.00", "+10.00"},
		{"-101", "-101"},
		{"+1023", "+1,023"},
		{"1288.6111", "1,288.6111"},
		{"-866444.259", "-866,444.259"},
		{"+31415986.283", "+31,415,986.283"},
		{"2222555666", "2,222,555,666"},
	} {
		if got := Enhanced_comma(tt.a); got != tt.want {
			t.Errorf("Enhanced Comma(%s) should return %v: but it returned: %v", tt.a, tt.want, got)
		}
	}
}

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
