package site

import (
	"errors"
	"fmt"
	"html/template"

	"github.com/hsson/go-website/post"
)

// Config is used for generating the site.
type Config struct {
	StylesheetFiles     []string   `yaml:"-"`
	Posts               post.Posts `yaml:"-"`
	OutputDirectory     string     `yaml:"-"`
	BaseURL             string     `yaml:"base_url"`
	TitleSuffix         string     `yaml:"title_suffix"`
	DefaultDescription  string     `yaml:"default_description"`
	GoogleAnalyticsCode string     `yaml:"google_analytics_code"`
	PostTemplateName    string     `yaml:"post_template"`
	PostDirectory       string     `yaml:"post_directory"`
	Pages               []Page     `yaml:"pages"`
}

// Page represents one page on the website
type Page struct {
	Title        string `yaml:"title"`
	Description  string `yaml:"description"`
	TempalteName string `yaml:"template"`
	URL          string `yaml:"url"`
}

// NewGenerator will initialize a new site generator implementing
// the Generator interface
func NewGenerator(config *Config, templates *template.Template) (Generator, error) {
	err := validateConfig(config)
	if err != nil {
		return nil, err
	}
	err = validateTemplatesExist(config, templates)
	return siteGenerator{config: config, templates: templates}, err
}

func validateTemplatesExist(config *Config, tempaltes *template.Template) error {
	template := tempaltes.Lookup(config.PostTemplateName)
	if template == nil {
		return fmt.Errorf("post template \"%s\" does not exist", config.PostTemplateName)
	}
	for _, page := range config.Pages {
		template := tempaltes.Lookup(page.TempalteName)
		if template == nil {
			return fmt.Errorf("template for page \"%s\" does not exist: %s", page.Title, page.TempalteName)
		}
	}
	return nil
}

func validateConfig(config *Config) error {
	if config.OutputDirectory == "" {
		return errors.New("no output directory specified")
	}
	if config.BaseURL == "" {
		return errors.New("a base url is required")
	}
	return nil
}
