package post

import (
	"fmt"
	"html/template"
	"time"
)

// Post represents a blog post with both metadata and content
type Post struct {
	Title    string        `yaml:"title"`
	Slug     string        `yaml:"slug"`
	Created  time.Time     `yaml:"created"`
	Updated  time.Time     `yaml:"updated"`
	Location string        `yaml:"location"`
	Author   string        `yaml:"author"`
	Content  template.HTML `yaml:"-"`
}

const timeFormat = "2006-01-02 15:04 (UTC-0700)"

// Posts is a sortable collection of Post
type Posts []*Post

// Len is the number of posts
func (p Posts) Len() int { return len(p) }

// Swap will swap the position of two posts
func (p Posts) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// ByCreationDate implements sort.Interface and sorts posts by their
// creation date
type ByCreationDate struct{ Posts }

// Less will determine which of two posts should be first
func (by ByCreationDate) Less(i, j int) bool {
	return by.Posts[i].Created.Before(by.Posts[j].Created)
}

func (p Post) String() string {
	return fmt.Sprintf("%s — %s", p.Created.Format(dateFormat), p.Title)
}

// FormattedDate gives the creation data of the post with some nice formatting
func (p Post) FormattedDate() string {
	return p.Created.Format(timeFormat)
}
