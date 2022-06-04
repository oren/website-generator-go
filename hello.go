package main

import (
	"github.com/shurcooL/github_flavored_markdown"
	// "github.com/shurcooL/github_flavored_markdown/gfmstyle"
	"os"
	"io/fs"
	"path/filepath"
	"fmt"
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

func main() {
	err := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() {
			fmt.Printf("visited file or dir: %q\n", path)
			/// convertFile(path)
			return nil
		}
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", ".", err)
		return
	}
}
