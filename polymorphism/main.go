// Polymorphism  is the ability to present the same interface for differing
// underlying forms (data types). Go allows us to use polymorphishm through
// "interface"
package main

import "fmt"

// Animal represents an animal
type Animal interface {
	sound() string
}

// Cat is a cat
type Cat struct {
	name string
}

func (cat Cat) sound() string {
	return "meow!"
}

// Dog is a dog
type Dog struct {
	name string
}

func (dog Dog) sound() string {
	return "woof!"
}

func main() {
	dog := Dog{name: "Puppy"}
	cat := Cat{name: "Kitty"}

	animals := []Animal{dog, cat}
	for _, animal := range animals {
		fmt.Println(animal.sound())
	}
}
