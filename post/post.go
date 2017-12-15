package post

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"sort"
	"time"

	slugify "github.com/mozillazg/go-slugify"
	blackfriday "gopkg.in/russross/blackfriday.v2"

	"gopkg.in/yaml.v2"
)

// FileEnding is the file ending that parsed posts have
const FileEnding = "md"
const postPattern = "*." + FileEnding
const dateFormat = "2006-01-02"

const metadataTemplate = `---
title: %s
slug: %s
created: %s
updated: %s
location: %s
author: %s
---`

type postChannelRes struct {
	sourceFile string
	post       *Post
	err        error
}

// LoadAll with load all the files in the specified directory
// and return a slice of all posts
func LoadAll(inputDir string) (result []*Post, err error) {
	fmt.Printf("Loading posts from %s\n", inputDir)

	postFiles, err := filepath.Glob(filepath.Join(inputDir, postPattern))
	if err != nil {
		return result, err
	}
	fmt.Printf("Found %d post(s)\n", len(postFiles))

	resChannel := make(chan postChannelRes)
	for _, postFile := range postFiles {
		go processPostRoutine(postFile, resChannel)
	}
	for i := 0; i < len(postFiles); i++ {
		res := <-resChannel
		if res.err != nil {
			fmt.Printf("Failed parsing of \"%s\"\n", res.sourceFile)
			return result, res.err
		}
		result = append(result, res.post)
	}
	sort.Sort(ByCreationDate{result})
	return result, err
}

// ParseFile will read a file with mixed metadata and content,
// parse it and then return a Post object
func ParseFile(inFile string) (*Post, error) {
	file, err := os.Open(inFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	inMeta := false
	var metadata, content bytes.Buffer
	for scanner.Scan() {
		line := scanner.Text()
		if line == "---" && !inMeta {
			inMeta = true
		} else if !inMeta {
			content.WriteString(line)
			content.WriteString("\n")
		} else if inMeta && line != "---" {
			metadata.WriteString(line)
			metadata.WriteString("\n")
		} else if inMeta && line == "---" {
			inMeta = false
		}
	}

	post := new(Post)
	err = yaml.Unmarshal(metadata.Bytes(), post)
	if err != nil {
		return post, err
	}
	formatted := blackfriday.Run(content.Bytes())
	post.Content = template.HTML(formatted)
	if err := scanner.Err(); err != nil {
		return post, err
	}
	return post, nil
}

// CreateMetadataString creates a correct metadata string that can be placed on the
// top of a post file.
func CreateMetadataString(title, slug, location, author string, created time.Time) string {
	return fmt.Sprintf(metadataTemplate, title, slug, created, created, location, author)
}

// SlugFromPost will create a post slug based on its title
func SlugFromPost(title string, creationDate time.Time) (string, error) {
	if title == "" {
		return "", errors.New("can't create slug from an empty title")
	}
	datePrefix := creationDate.Format(dateFormat)
	slug := slugify.Slugify(title)
	return fmt.Sprintf("%s-%s", datePrefix, slug), nil
}

func processPostRoutine(inputFile string, resChannel chan postChannelRes) {
	post, err := ParseFile(inputFile)
	resChannel <- postChannelRes{inputFile, post, err}
}
