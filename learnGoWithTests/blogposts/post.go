package blogposts

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
)

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	title := readMetaLine(titleSeparator)
	description := readMetaLine(descriptionSeparator)
	tags := strings.Split(readMetaLine(tagsSeparator), ", ")
	body := readBody(scanner)

	return Post{Title: title, Description: description, Tags: tags, Body: body}, nil
}

func readBody(scanner *bufio.Scanner) string {
	var text string
	scanner.Scan()
	for scanner.Scan() {
		text += fmt.Sprintf("%s\n", scanner.Text())
	}
	text = strings.TrimSuffix(text, "\n")
	return text
}
