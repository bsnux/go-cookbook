package main

import "fmt"

// Persona uses first and last name
type Persona struct {
	first string
	last  string
}

// changeName will change `first` using a pointer receiver
// There are two reasons to use a pointer receiver:
// 	1) The method can modify the value that its receiver points to
//  2) To avoid copying the value on each method call. This can be more efficient
//     if the receiver is a large struct, for example.
func (p *Persona) changeName(newName string) {
	p.first = newName
}

func main() {
	// Declaring a pointer to `int`
	var p *int

	i := 7
	// `&` generates a pointer to its operand
	p = &i

	fmt.Printf("i= %d\n", i)
	fmt.Printf("p= %d\n", *p)

	i = 8
	fmt.Printf("i= %d\n", i)
	fmt.Printf("p= %d\n", *p)

	byVal(i)
	fmt.Printf("i by val: %d\n", i)

	byRef(&i)
	fmt.Printf("i by ref: %d\n", i)

	// Returning a pointer to the struct value.
	persona := &Persona{"Foo", "Bar"}
	fmt.Println(persona.first)

	// Structs and pointers
	persona.changeName("New name")
	fmt.Println(persona.first)
}

func byVal(a int) {
	a = 2
}

func byRef(a *int) {
	*a = 3
}
