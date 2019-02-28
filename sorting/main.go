// How to sort data
//
// sort.Sort function is an implementation of the quicksort algorithm.
// This means at best a call to sort.Sort is O(n*log(n))
package main

import (
	"fmt"
	"sort"
)

func main() {
	// Sort methods are specific to the builtin type
	strs := []string{"xy", "ab", "cd"}
	sort.Strings(strs)
	fmt.Println(strs)

	numbers := []int{9, 1, 4}
	sort.Ints(numbers)
	fmt.Println(numbers)

	m := map[string]int{"Bob": 25, "Alice": 23, "Eve": 20, "Charlie": 21}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, m[k])
	}
}
