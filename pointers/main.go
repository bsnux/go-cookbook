package main

import "fmt"

// Persona uses first and last name
type Persona struct {
	first string
	last  string
}

// changeName will change `first` using a pointer receiver
// There are two reasons to use a pointer receiver:
//  1. The method can modify the value that its receiver points to
//  2. To avoid copying the value on each method call. This can be more efficient if the receiver is a large struct
func (p *Persona) changeName(newName string) {
	p.first = newName
}

// changeName2 won't change original struct
func (p Persona) changeName2(newName string) {
	p.first = newName
}

func main() {
	// Declaring a pointer to `int`.
	// A pointer is represented using `*` followed by the type of the stored value
	// `*` is also used to dereference pointer variables, which gives us access to the value the pointer points to
	var p *int

	i := 7
	// `&` generates a pointer to its operand
	// We use the `&` operator to find the address of a variable
	// `&i` returns a pointer to an int because `i` is an int
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
	fmt.Println((*persona).first)

	// You can use `&` with same effect that lines above. Go figure that out for you
	persona2 := Persona{"Foo2", "Bar2"}
	fmt.Println(persona2.first)
	// Next line won't compile because we used `Persona` and not `&Persona`
	// fmt.Println((*persona2).first)

	// Structs and pointers
	persona.changeName("New name")
	fmt.Println(persona.first)

	persona2.changeName("New name 2")
	fmt.Println(persona2.first)

	persona.changeName2("The new name")
	fmt.Println(persona.first)
}

func byVal(a int) {
	a = 2
}

func byRef(a *int) {
	*a = 3
}
