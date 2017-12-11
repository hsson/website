package util

import (
	"io/ioutil"

	"github.com/hsson/go-website/site"
	yaml "gopkg.in/yaml.v2"
)

// LoadConfig loads the config at the specified path into a struct
// and returns it
func LoadConfig(configPath string) (*site.Config, error) {
	config := new(site.Config)
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	err = yaml.UnmarshalStrict(data, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
