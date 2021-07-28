// Exercise 1.1 Modify the echo program to also print os.Args[0], the name of the command that invoked it.

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 4.
//!+

package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Printf(s)
	fmt.Printf("The command that invoked os.Args is %v", os.Args[0])
}

//!-
