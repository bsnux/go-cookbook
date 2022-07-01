package main

import "fmt"

func main() {
	var slice []int
	var array [10]int
	dict := make(map[string]int)

	array[0] = 1
	fmt.Println(array)

	slice = append(slice, 2)
	fmt.Println(slice)

	dict["hello"] = 1
	fmt.Println(dict)
}
