package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/hsson/go-website/post"
	"github.com/hsson/go-website/util"
)

func createNewPost() {
	fmt.Println("Creating new post")
	newPost := createPostBase()
	postMetadata := post.CreateMetadataString(newPost.Title, newPost.Slug, newPost.Location, newPost.Author, newPost.Created)
	fmt.Println(postMetadata)

	fileName := fmt.Sprintf("%s.%s", newPost.Slug, post.FileEnding)
	newPostFilename := filepath.Join(postsIn, fileName)
	err := ioutil.WriteFile(newPostFilename, []byte(postMetadata), 0666)
	util.MaybeExitWithError(err)
	err = editPost(newPostFilename)
	util.MaybeExitWithError(err)
	fmt.Printf("Created %s\n", newPostFilename)
}

func createPostBase() post.Post {
	postTitle, err := parseTitle(readLineWithPrompt(os.Stdin, "Enter title: "))
	util.MaybeExitWithError(err)
	postAuthor := parseWithDefault(readLineWithPrompt(os.Stdin, "Enter author: "), defaultAuthor)
	postLocation := parseWithDefault(readLineWithPrompt(os.Stdin, "Enter location: "), defaultLocation)
	creationDate := time.Now()
	postSlug, err := post.SlugFromPost(postTitle, creationDate)
	util.MaybeExitWithError(err)
	return post.Post{
		Title:    postTitle,
		Slug:     postSlug,
		Created:  creationDate,
		Updated:  creationDate,
		Location: postLocation,
		Author:   postAuthor,
	}
}

func editPost(postFile string) error {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = fallbackEditor
	}
	cmd := exec.Command(editor, postFile)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		return err
	}
	return cmd.Wait()
}

func parseTitle(rawTitle string) (string, error) {
	if rawTitle == "" {
		return "", errors.New("post title can't be empty")
	}
	return rawTitle, nil
}

func parseWithDefault(inputValue, defaultValue string) string {
	if inputValue == "" {
		return defaultValue
	}
	return inputValue
}

func readLine(in io.Reader) string {
	reader := bufio.NewReader(in)
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(text)
}

func readLineWithPrompt(in io.Reader, prompt string) string {
	fmt.Print(prompt)
	return readLine(in)
}
