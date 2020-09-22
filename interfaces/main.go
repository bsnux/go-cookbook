package main

import (
	"fmt"
	"regexp"
	"strings"
)

type Post struct {
	Title  string
	Author string
}

type Article struct {
	Title    string
	Author   string
	Category string
}

type Linker interface {
	Link() string
}

func normalize(field string) string {
	regex := regexp.MustCompile(`\s+`)
	return regex.ReplaceAllString(strings.ToLower(field), "-")
}

// Link implements Link method in interface so Post satisfies Linker interface
func (p Post) Link() string {
	return fmt.Sprintf("%s-%s\n", normalize(p.Title), normalize(p.Author))
}

// Link implements Link method in interface so Article satisfies Linker interface
func (a Article) Link() string {
	return fmt.Sprintf("%s-%s-%s\n", normalize(a.Title), normalize(a.Author), normalize(a.Category))
}

// PermaLink calls to Link method for every struct passed as parameter
// which satisfies the Linker interface
func permaLink(l Linker) {
	fmt.Println(l.Link())
}

func main() {
	p := Post{Title: "The post", Author: "The author"}
	a := Article{Title: "The article", Author: "The author", Category: "The Category"}

	permaLink(p)
	permaLink(a)
}
