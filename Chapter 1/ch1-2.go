package main

import (
	"bufio"
	"fmt"
	"os"
)

const question12 = `Exercise 1.2 Modify the echo program to print the index and
// value of each of its arguements, one per line.`

const question13 = `Excercise 1.3 Experiment to measure the difference in running time between our 
// potentially inefficent version and the one that uses strings.Join (Section 1.6 illustrates part of the
// time package, and secion 11.4 shows how to write benchmark test for systematic performace evaluation.)`

func main() {
	fmt.Println(question12)
	problem11()
	problem12()
	problem14()
}

func problem11() {
	fmt.Println("Starting problem 1.1")
	var s string
	for i := 0; i < len(os.Args); i++ {
		s += fmt.Sprintf("%v \n", os.Args[i])
	}
	fmt.Printf(s)
}

func problem12() {
	fmt.Println("starting problem1.2")
	var s string
	for i := 0; i < len(os.Args); i++ {
		s += fmt.Sprintf("Argument %v is %v \n", i, os.Args[i])
	}
	fmt.Printf(s)
}

const question14 = `Modiffy dup2 to print the names of all files in which each duplicated line occures`

func problem14() {
	fmt.Println("Starting problem 1.4")
	counts := make(map[string]map[string]int)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, filecnt := range counts {
		if len(filecnt) > 1 {
			var present []string
			for k := range filecnt {
				present = append(present, k)
			}
			fmt.Printf("The line %q is present in %v files \n", line, present)

		} else {
			for k, v := range filecnt {
				if v > 1 {
					fmt.Printf("The line %q is present in %q files \n", line, k)
				}
			}
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if len(counts[input.Text()]) == 0 {
			counts[input.Text()] = make(map[string]int)
		}
		counts[input.Text()][f.Name()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
