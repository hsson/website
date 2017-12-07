package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"github.com/hsson/go-website/post"
	"github.com/hsson/go-website/style"
)

const banner = ` _             _                                                               
| |           | |                                                              
| |__    __ _ | | __  __ _  _ __   ___  ___   ___   _ __     __  __ _   _  ____
| '_ \  / _` + "`" + ` || |/ / / _` + "`" + ` || '_ \ / __|/ __| / _ \ | '_ \    \ \/ /| | | ||_  /
| | | || (_| ||   < | (_| || | | |\__ \\__ \| (_) || | | | _  >  < | |_| | / / 
|_| |_| \__,_||_|\_\ \__,_||_| |_||___/|___/ \___/ |_| |_|(_)/_/\_\ \__, |/___|
                                                                     __/ |     
                                                                    |___/      `

var (
	outputDir = flag.String("out", "build", "Specify the output directory where generated files will be placed")

	styleIn  = filepath.Join("resources", "styles")
	styleOut = filepath.Join("assets", "css")

	postsIn = filepath.Join("resources", "posts")
)

func main() {
	flag.Parse()
	fmt.Println(banner)
	fmt.Printf("Generated files will be put in \"%s\"...\n", *outputDir)
	sheets, err := style.Generate(styleIn, *outputDir, styleOut)
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
