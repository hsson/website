package main

import (
	"fmt"
	"html/template"
	"path/filepath"

	"github.com/hsson/go-website/post"
	"github.com/hsson/go-website/site"
	"github.com/hsson/go-website/style"
	"github.com/hsson/go-website/util"
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

	fmt.Println("Created posts:")
	for _, post := range posts {
		fmt.Println(post)
	}

	config, err := util.LoadConfig(*flagConfigFile)
	if err != nil {
		panic(err)
	}
	config.StylesheetFiles = sheets
	config.Posts = posts
	config.OutputDirectory = *flagOutputDir

	fmt.Printf("Loading templates form %s\n", templatesIn)
	templates, err := template.ParseGlob(filepath.Join(templatesIn, "*.tmpl"))
	if err != nil {
		panic(err)
	}
	generator, err := site.NewGenerator(config, templates)
	if err != nil {
		panic(err)
	}
	err = generator.Generate(*flagOutputDir)
	if err != nil {
		panic(err)
	}
	fmt.Println("Website generation complete")
}
