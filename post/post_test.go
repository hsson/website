package post

import (
	"fmt"
	"testing"
	"time"
)

func TestMetadata(t *testing.T) {
	title := "Hej på dig"
	slug, _ := SlugFromTitle(title)
	theTime, _ := time.Parse(time.RFC3339, "2017-12-08T16:53:23.236728-08:00")
	author := "John Doe"
	location := "Göteborg"
	answer := `---
title: Hej på dig
slug: 2017-12-08-hej-pa-dig
created: 2017-12-08 16:53:23.236728 -0800 PST
updated: 2017-12-08 16:53:23.236728 -0800 PST
location: Göteborg
author: John Doe
---`
	result := CreateMetadataString(title, slug, location, author, theTime)
	if result != answer {
		t.Errorf("Expected:\n%s\nGot:\n%s\n", answer, result)
	}
}

func TestSlug(t *testing.T) {
	title := "Hej på dig"
	datePrefix := time.Now().Format(dateFormat)
	slug, err := SlugFromTitle(title)
	if err != nil {
		t.Errorf("The title should be okay")
	}
	if slug != fmt.Sprintf("%s-%s", datePrefix, "hej-pa-dig") {
		t.Errorf("Got bad slug '%s'", slug)
	}
}

func TestSlugEmptyTitle(t *testing.T) {
	emptyTitle := ""
	_, err := SlugFromTitle(emptyTitle)
	if err == nil {
		t.Errorf("Should not allow slugs from empty titles")
	}
}
