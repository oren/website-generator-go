package main

import (
	"github.com/shurcooL/github_flavored_markdown"
	// "github.com/shurcooL/github_flavored_markdown/gfmstyle"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func convertFile() {
	// convert README.md to index.html
	dat, err := os.ReadFile("README.md")
	check(err)
	output := github_flavored_markdown.Markdown(dat)
	err = os.WriteFile("index.html", output, 0644)
	check(err)
}

func convertFolder() {
	convertFile()
	// for each folder, call ConvertFolder
}

func main() {
	convertFolder()
}
