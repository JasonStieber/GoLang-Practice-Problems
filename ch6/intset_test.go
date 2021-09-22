// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import (
	"fmt"
	"testing"
)

func Example_one() {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func Example_two() {
	var x IntSet
	x.Add(1, 144, 9, 42)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	//!-note

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}

func TestLen(t *testing.T) {
	type test = struct {
		n    IntSet
		want int
	}
	var u, v, w, x, y, z IntSet
	u.Add(2, 7777, 69)
	w.Add(9999999)
	x.Add(1, 23)
	y.Add(22, 22, 22, 22, 21)
	z.Add(877, 997, 0, 28468421561)

	for _, tt := range []test{
		{v, 0},
		{u, 3},
		{w, 1},
		{x, 2},
		{y, 2},
		{z, 4},
	} {
		if got := tt.n.Len(); got != tt.want {
			t.Fatalf("Len() should be %v, but is: %v", tt.want, got)
		}
	}
}

func TestRemove(t *testing.T) {
	type test = struct {
		s    IntSet
		v    int
		want string
	}
	var u, v, w, x, y, z IntSet
	u.Add(2, 7777, 69)
	w.Add(9999999)
	x.Add(1, 23)
	y.Add(2, 20, 22, 42, 21)
	z.Add(877, 997, 0, 28468421561)

	for _, tt := range []test{
		{v, 3, "{}"},
		{u, 7777, "{2 69}"},
		{w, 9999999, "{}"},
		{x, 1, "{23}"},
		{y, 22, "{2 20 21 42}"},
		{z, 4, "{0 877 997 28468421561}"},
	} {
		or := tt.s.String()
		tt.s.Remove(tt.v)
		if got := tt.s.String(); got != tt.want {
			t.Fatalf("Remove %v from %v should be %v not %v", tt.v, or, tt.want, got)
		}
	}
}

func TestCopy(t *testing.T) {
	type test struct {
		s    IntSet
		want string
	}

	var u, w, x, y, z IntSet
	u.Add(2, 7777, 69)
	w.Add(9999999)
	x.Add(1, 23)
	y.Add(2, 20, 22, 42, 21)
	z.Add(877, 997, 0, 28468421561)

	for _, tt := range []test{
		{u, "{2 69 7777}"},
		{w, "{9999999}"},
		{x, "{1 23}"},
		{y, "{2 20 21 22 42}"},
		{z, "{0 877 997 28468421561}"},
	} {
		got := tt.s.Copy().String()
		tt.s.Add(3, 54, 6, 7, 8)
		if got != tt.want {
			t.Fatalf("Copy should have returned IntSet %v, but it returned %v", tt.want, got)
		}
		tt.s.Clear()
		if got != tt.want {
			t.Fatalf("Copy should have returned IntSet %v, but it returned %v", tt.want, got)
		}
	}
}
