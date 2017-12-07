package post

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"gopkg.in/yaml.v2"
)

const postPattern = "*.md"

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
		fmt.Printf("Created post \"%s\"\n", res.post.Title)
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
	post.Content = content.String()
	if err := scanner.Err(); err != nil {
		return post, err
	}
	return post, nil
}

func processPostRoutine(inputFile string, resChannel chan postChannelRes) {
	post, err := ParseFile(inputFile)
	resChannel <- postChannelRes{inputFile, post, err}
}
