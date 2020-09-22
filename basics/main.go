package main

import "fmt"

func main() {
	hello()
	arrays()
	slices()
	maps()
	variadicFunc()
}

func hello() {
	fmt.Println("Hello")
	print("Hello World\n")

	// `rune` is an alias for int32. It is an UTF-8 encoded code point
	s := "hello"
	cad := []rune(s)
	cad[0] = 'c'
	fmt.Println(string(cad))
}

func arrays() {
	// static array
	var arr [10]int

	arr[0] = 1
	arr[1] = 2

	fmt.Printf("The first element is %d\n", arr[0])
	fmt.Println(arr)

	// initializating array
	var arrInt = [5]int{1, 2, 3, 4, 5}
	fmt.Println(len(arrInt))
	fmt.Println(arrInt)
}

func slices() {
	// A slice is similar to an array, but it can grow when new elements
	// are added. A slice always refers to an underlying array. What makes
	// slices different from arrays is that a slice is a pointer to an array.
	// Slices are reference types.
	// Slices can be created with the built-in make function; this is how you
	// create dynamically-sized arrays
	ss := make([]int, 10)
	ss[8] = 2
	fmt.Println(ss)
	fmt.Printf("len is %d\n", len(ss))
	fmt.Printf("capicity is %d\n", cap(ss))

	// Arrays: `...` notation specifies a length equal to the number of elements
	// in the literal
	a := [...]int{1, 2, 3, 4, 5}
	newSlice := a[3:4]
	fmt.Println(newSlice)
	fmt.Printf("len is %d\n", len(newSlice))
	fmt.Printf("capicity is %d\n", cap(newSlice))

	// Creating slices without make is also valid
	var dynArr []int
	dynArr = append(dynArr, 3)
	dynArr = append(dynArr, 4)
	fmt.Println(dynArr)
}

func maps() {
	testMap := make(map[string]int)

	testMap["hello"] = 10
	testMap["bye"] = 20

	fmt.Println(testMap)
	fmt.Println(testMap["hello"])

	for k, v := range testMap {
		fmt.Printf("key is %s, val is %d\n", k, v)
	}

	_, exists := testMap["hi"]
	if exists {
		fmt.Println(testMap["hi"])
	}

	monthdays := map[string]int{
		"Jan": 31, "Feb": 28, "Mar": 31,
		"Apr": 30, "May": 31, "Jun": 30,
		"Jul": 31, "Aug": 31, "Sep": 30,
		"Oct": 31, "Nov": 30, "Dec": 31,
	}
	fmt.Println(monthdays)

	year := 0
	for _, days := range monthdays {
		year += days
	}
	fmt.Printf("Numbers of days in a year: %d\n", year)
}

func printNumbers(nums ...int) {
	fmt.Println(nums)
}

func variadicFunc() {
	printNumbers(1, 2, 3)
	printNumbers(1, 2)
}
