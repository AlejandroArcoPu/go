package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Item struct {
	Title, URL string
}

func main() {
	const blob = `
[
	{"Title": "Google", "URL": "http://google.com"},
	{"Title": "Facebook", "URL": "http://facebook.com"}
]
`
	var items []*Item
	json.NewDecoder(strings.NewReader(blob)).Decode(&items)
	for _, v := range items {
		fmt.Printf("Title: %v URL: %v", v.Title, v.URL)
	}
}
