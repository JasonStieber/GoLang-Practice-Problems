package main

import (
	"fmt"
	"os"
)

const question12 = `Exercise 1.2 Modify the echo program to print the index and
// value of each of its arguements, one per line.`

func main() {
	fmt.Println(question12)
	problem11()
	problem12()
}

func problem11() {
	var s string
	for i := 0; i < len(os.Args); i++ {
		s += fmt.Sprintf("%v \n", os.Args[i])
	}
	fmt.Printf(s)
}

func problem12() {
	var s string
	for i := 0; i < len(os.Args); i++ {
		s += fmt.Sprintf("Argument %v is %v \n", i, os.Args[i])
	}
	fmt.Printf(s)
}
