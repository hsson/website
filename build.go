package main

import (
	"fmt"

	"github.com/hsson/go-website/post"
	"github.com/hsson/go-website/style"
)

func buildSite() {
	fmt.Printf("Generated files will be put in \"%s\"...\n", *flagOutputDir)
	sheets, err := style.Generate(styleIn, *flagOutputDir, styleOut)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created sheets: %v\n", sheets)

	posts, err := post.LoadAll(postsIn)
	if err != nil {
		panic(err)
	}

	for _, post := range posts {
		fmt.Println(post.Content)
	}
}
