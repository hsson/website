package site

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/hsson/go-website/post"
)

// Generator is used to generate the website based on some given config
type Generator interface {
	Generate(outputDir string) error
}

type pageData struct {
	Page
	Config
	Post post.Post
}

type siteGenerator struct {
	config    *Config
	templates *template.Template
}

const fileEnding = "html"

// Generate will build the website by generating all nessecary files
// based on the provided config. All files will be put in the specified
// outputDir
func (g siteGenerator) Generate(outputDir string) error {

	resultChannel := make(chan buildChannelRes)
	launched := 0
	// Create all individual post pages
	postOutDir := g.config.PostDirectory
	os.MkdirAll(filepath.Join(outputDir, postOutDir), os.ModePerm)
	for _, p := range g.config.Posts {
		filename := fmt.Sprintf("/%s/%s.%s", postOutDir, p.Slug, fileEnding)
		page := &Page{
			Title:        p.Title,
			Description:  "A post", // TODO: Fix
			TempalteName: g.config.PostTemplateName,
			URL:          filename,
		}

		outFile := filepath.Join(outputDir, postOutDir, fmt.Sprintf("%s.%s", p.Slug, fileEnding))

		template := g.templates.Lookup(g.config.PostTemplateName)

		pageData := &pageData{
			Page:   *page,
			Config: *g.config,
			Post:   *p,
		}
		go buildPageRoutine(outFile, pageData, template, resultChannel)
		launched++
	}

	// Create the other pages
	for _, p := range g.config.Pages {
		filename := fmt.Sprintf("%s.%s", p.TempalteName, fileEnding)
		outFile := filepath.Join(outputDir, filename)
		template := g.templates.Lookup(p.TempalteName)
		pageData := &pageData{
			Page:   p,
			Config: *g.config,
		}
		go buildPageRoutine(outFile, pageData, template, resultChannel)
		launched++
	}

	// Wait for all pages to build
	for i := 0; i < launched; i++ {
		res := <-resultChannel
		if res.err != nil {
			fmt.Printf("Failed to create %s\n", res.targetFile)
			return res.err
		}
	}
	return nil
}

type buildChannelRes struct {
	targetFile string
	err        error
}

func buildPageRoutine(outputFile string, pageData *pageData, template *template.Template, ch chan buildChannelRes) {
	res := buildChannelRes{err: nil, targetFile: outputFile}

	file, err := os.Create(outputFile)
	defer file.Close()
	if err != nil {
		res.err = err
		ch <- res
		return
	}

	err = template.Execute(file, pageData)
	if err != nil {
		res.err = err
	}
	ch <- res
}
