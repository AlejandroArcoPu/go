package main

import (
	"fmt"
	"os"

	"github.com/AlejandroArcoPu/go/learnGoWithTests/blogposts"
)

func main() {
	root := "./posts"
	fileSystem := os.DirFS(root)
	posts, err := blogposts.NewPostsFromFS(fileSystem)
	if err != nil {
		fmt.Println("an error has occured with your markdown: ", err.Error())
	}
	fmt.Println(posts)
}
