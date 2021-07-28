// Exercise 1.1 Modify the echo program to also print os.Args[0], the name of the command that invoked it.

// Origional example

package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	fmt.Println("are we seeing this?")
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Printf("This is the varabile we are working with %v", os)
	fmt.Println(s)
}

//!-
