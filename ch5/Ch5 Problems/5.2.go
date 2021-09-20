package main

// write a function that creates a map with keys of the HTML element names and the value set to the
// number of times each element is seen

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	n, err := html.Parse(strings.NewReader(rawHTML))
	if err != nil {
		panic(err)
	}
	m := make(map[string]int)
	outline(m, n)
	fmt.Println(m)
}

// _ io.Reader  := http.Response{}.Body

// write a fetch function to grab html data
func fetch(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Get request from %v resulted in a %v status code", url, resp.StatusCode)
	}
	n, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to Parse data from URL: %v", url)
	}
	return n, nil
}

func countElements(url string) map[string]int {
	// If i'm taking the data from standardin why does countElemennts needs to have a param passed too it
	// and there is no variable name, is this just the standard in?? if its not used why isnt' there an error
	var eleCnt map[string]int
	n, err := fetch(url)
	if err != nil {
		fmt.Printf(url, "Count Elements: %v\n", err)
	}
	outline(eleCnt, n)
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
