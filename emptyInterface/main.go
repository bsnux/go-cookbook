package main

import "fmt"

// printVal will print different message based on type. Using empty interface
// as param for using different types
func printVal(v interface{}) {
	switch v := v.(type) {
	case string:
		fmt.Printf("%v is a string\n", v)
	case int:
		fmt.Printf("%v is an int\n", v)
	default:
		fmt.Printf("The type of v is unknown\n")
	}
}

func main() {
	v := 10
	printVal(v)

	cad := "hello"
	printVal(cad)
}
