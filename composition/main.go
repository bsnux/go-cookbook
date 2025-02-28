// Go doesn't support inheritance but allows us to use composition for similar
// purposes
// Composition over inheritance in object-oriented programming (OOP) is the
// principle that classes should achieve polymorphic behavior and code
// reuse by their composition (by containing instances of other classes that
// implement the desired functionality) rather than inheritance
// from a base or parent class
package main

import "fmt"

type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type post struct {
	title   string
	content string
	author
}

func (p post) info() {
	fmt.Printf("Title: %s\tAuthor: %s\tBio: %s", p.title, p.fullName(), p.bio)
}

type blog struct {
	posts []post
}

func (b blog) contents() {
	fmt.Println("Blog content:")
	for _, post := range b.posts {
		post.info()
		fmt.Println()
	}
}

func main() {
	author := author{"Name", "Last name", "This is the main author"}
	post1 := post{"First post", "This is the content for first post", author}
	post2 := post{"Second post", "This is the content for second post", author}

	b := blog{
		posts: []post{post1, post2},
	}

	b.contents()

	// Calling to property and method in `Author` from `Post`
	fmt.Println(post1.lastName)
	fmt.Println(post1.fullName())
}
