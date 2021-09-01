package Ch5_Problems

// write a function that creates a map with keys of the HTML element names and the value set to the
// number of times each element is seen

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

// write a fetch function to grab html data
func fetch() {

}
func countElements(*html.Node) map[string]int {
	var eleCnt map[string]int
	doc, err = html.Parse(os.Stdin)
	if err != nil {
		fmt.fprintf(os.Stderr, "Count Elements: %v\n", err)
		os.Exit(1)
	}
	outline(eleCnt, doc)
	fmt.Printf("The element tag count is: %v", eleCnt)
	return eleCnt
}

func outline(eleCnt map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		eleCnt[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(eleCnt, c)
	}
}
