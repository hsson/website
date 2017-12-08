package main

import (
	"strings"
	"testing"
)

func TestParseShouldDefault(t *testing.T) {
	defaultValue := "TheDefault"
	inputValue := ""
	result := parseWithDefault(inputValue, defaultValue)

	if result != defaultValue {
		t.Errorf("The default value should have been used")
	}
}

func TestParseShouldNotDefault(t *testing.T) {
	defaultValue := "TheDefault"
	inputValue := "SomeValue"
	result := parseWithDefault(inputValue, defaultValue)
	if result != inputValue {
		t.Errorf("The value is not the same as the one entered")
	}
}

func TestReadNormalLine(t *testing.T) {
	inputLine := "some input line"
	mockReader := strings.NewReader(inputLine + "\n")
	result := readLine(mockReader)
	if result != inputLine {
		t.Errorf("The resulting string is not the one inputted")
	}
}

func TestReadNonEndingLine(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	inputLine := "some input line"
	mockReader := strings.NewReader(inputLine)
	readLine(mockReader)
}

func TestReadWithPrompt(t *testing.T) {
	inputLine := "some input line"
	mockReader := strings.NewReader(inputLine + "\n")
	result := readLineWithPrompt(mockReader, "some prompt")
	if result != inputLine {
		t.Errorf("Didn't get the correct input")
	}
}

func TestTitleValid(t *testing.T) {
	validTitle := "An awesome title"
	result, err := parseTitle(validTitle)
	if err != nil {
		t.Errorf("Title was incorrectly marked as invalid")
	}
	if result != validTitle {
		t.Errorf("Title was changed")
	}
}

func TestTitleEmpty(t *testing.T) {
	emptyTitle := ""
	_, err := parseTitle(emptyTitle)
	if err == nil {
		t.Errorf("An empty title is not allowed")
	}
}
