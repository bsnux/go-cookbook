package main

/**
* arrays, slices, and maps 101
 */

import (
	"fmt"
)

func main() {
	// Arrays
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{1, 2, 3, 4}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	// Slices
	// nil slice
	var slice1 []int
	slice2 := make([]string, 3)

	fmt.Println(slice1)
	fmt.Println(slice2)

	slice1 = append(slice1, 1)
	fmt.Println(slice1)

	// We don't use append here because it's not a nil slice, we already created the slice with lenght and capacity of 3
	slice2[0] = "bob"
	slice2[1] = "alice"
	slice2[2] = "frank"
	fmt.Println(slice2)

	// Maps
	// nil map
	var map1 map[string]int
	// initializing nil map
	map1 = make(map[string]int)
	map1["frank"] = 3
	fmt.Println(map1)

	map2 := make(map[string]int)
	map2["bob"] = 1
	map2["alice"] = 2
	fmt.Println(map2)
}
