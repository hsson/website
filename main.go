package main

import (
	"flag"
	"fmt"
	"path/filepath"
)

const banner = ` _             _                                                               
| |           | |                                                              
| |__    __ _ | | __  __ _  _ __   ___  ___   ___   _ __     __  __ _   _  ____
| '_ \  / _` + "`" + ` || |/ / / _` + "`" + ` || '_ \ / __|/ __| / _ \ | '_ \    \ \/ /| | | ||_  /
| | | || (_| ||   < | (_| || | | |\__ \\__ \| (_) || | | | _  >  < | |_| | / / 
|_| |_| \__,_||_|\_\ \__,_||_| |_||___/|___/ \___/ |_| |_|(_)/_/\_\ \__, |/___|
                                                                     __/ |     
                                                                    |___/      `

var (
	flagOutputDir  = flag.String("out", "build", "Specify the output directory where generated files will be placed")
	flagNewPost    = flag.Bool("new", false, "Specify this flag if you want to create a new post")
	flagBuildSite  = flag.Bool("build", false, "Specify this flag if you want to build the whole site")
	flagConfigFile = flag.String("config", "config.yaml", "Specify the config file to be used to generate the site")

	styleIn  = filepath.Join("resources", "styles")
	styleOut = filepath.Join("assets", "css")

	postsIn = filepath.Join("resources", "posts")

	templatesIn = filepath.Join("resources", "templates")

	defaultAction = buildSite
)

const (
	defaultAuthor   = "Alexander Håkansson"
	defaultLocation = "San Francisco"
	fallbackEditor  = "vim"
)

func main() {
	flag.Parse()
	fmt.Println(banner)

	if *flagNewPost {
		createNewPost()
	} else if *flagBuildSite {
		buildSite()
	} else {
		defaultAction()
	}
}
