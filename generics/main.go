package main

import "fmt"

// printThis shows how to use type params
func printThis[T any](value T) {
	fmt.Println(value)
}

// index returns the index in the slice for the given item
// we need to use `comparable` in order to be able to compare values
func index[T comparable](vs []T, t T) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// getKeys returns the keys of the map
// The value type doesn’t have any restraints but the key type should always
// satisfy the comparable constraint.
func getKeys[K comparable, V any](m map[K]V) []K {
	key := make([]K, len(m))
	i := 0
	for k, _ := range m {
		key[i] = k
		i++
	}
	return key
}

func main() {
	printThis(34)
	printThis(2.50)
	printThis("this value")

	keys := getKeys(map[string]int{"one": 1, "two": 2})
	fmt.Println(keys)

	keysInt := getKeys(map[int]int{1: 1, 2: 2})
	fmt.Println(keysInt)


  arrStr := []string{"one", "two", "three"}
	arrInt := []int{1, 3, 2}
  idxStr := index(arrStr, "three")
  idxInt := index(arrInt, 3)
  fmt.Println(idxStr)
  fmt.Println(idxInt)
}
