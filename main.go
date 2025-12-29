package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gomarkdown/markdown"
)

const baseURL = "https://oren.github.io"

// convertNavbarURLs converts relative URLs in markdown links to absolute URLs
func convertNavbarURLs(mdContent string) string {
	// Match markdown links: [text](url)
	re := regexp.MustCompile(`\[([^\]]+)\]\(([^)]+)\)`)

	return re.ReplaceAllStringFunc(mdContent, func(match string) string {
		parts := re.FindStringSubmatch(match)
		if len(parts) != 3 {
			return match
		}
		text := parts[1]
		url := parts[2]

		// Convert relative URLs to absolute
		if url == "/" {
			url = baseURL
		} else if strings.HasPrefix(url, "/") {
			// /path -> baseURL/path (remove leading slash since baseURL doesn't have trailing)
			url = baseURL + url
		} else if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			// relative path like "cook/" -> baseURL/cook/
			url = baseURL + "/" + url
		}

		return fmt.Sprintf("[%s](%s)", text, url)
	})
}

func processDirectory(dir string) {
	readmePath := filepath.Join(dir, "README.md")
	if _, err := os.Stat(readmePath); os.IsNotExist(err) {
		return // Skip if no README.md
	}

	// fmt.Printf("Processing %s...\n", readmePath)

	// Read the README.md file
	md, err := ioutil.ReadFile(readmePath)
	if err != nil {
		fmt.Printf("Error reading %s: %v\n", readmePath, err)
		return
	}

	// Convert markdown to HTML
	html := markdown.ToHTML(md, nil, nil)

	// Check for _navbar.md and convert it
	navbarHTML := ""
	navbarPath := filepath.Join(dir, "_navbar.md")
	if _, err := os.Stat(navbarPath); err == nil {
		navbarMd, err := ioutil.ReadFile(navbarPath)
		if err == nil {
			// Convert relative URLs to absolute URLs
			navbarMdConverted := convertNavbarURLs(string(navbarMd))
			navbarHTML = string(markdown.ToHTML([]byte(navbarMdConverted), nil, nil))
		}
	}

	// Create the HTML content
	htmlContent := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>README</title>
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/github-markdown-css/5.5.1/github-markdown.min.css">
<style>
	body {
		font-size: 1.5em;
	}
	.markdown-body {
		box-sizing: border-box;
		min-width: 200px;
		max-width: 980px;
		margin: 0 auto;
		padding: 45px;
		background-color: white;
		color: #333;
		font-size: 0.8em;
	}

	@media (max-width: 767px) {
		.markdown-body {
			padding: 15px;
		}
	}
</style>
</head>
<body>
%s
<article class="markdown-body">
%s
</article>
</body>
</html>`, navbarHTML, html)

	// Write the index.html file
	indexPath := filepath.Join(dir, "index.html")
	err = ioutil.WriteFile(indexPath, []byte(htmlContent), 0644)
	if err != nil {
		fmt.Printf("Error writing %s: %v\n", indexPath, err)
		return
	}

	// fmt.Printf("Successfully converted %s to %s\n", readmePath, indexPath)
}

func main() {
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			processDirectory(path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error walking through directories:", err)
		os.Exit(1)
	}
}
