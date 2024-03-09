package main

import "fmt"

func main() {
	// Same as `var x int = 10`
	x := 10
	y := 20
	z := x + y
	fmt.Println(z)

	// Arrays
	// create and array by `var a [5]int` <- 5 is the length of the array, can't be changed
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)

	// Slices
	// create a slice by `var a []int` <- no length, can be changed
	b := []int{1, 2, 3, 4, 5}
	// append to a slice
	b = append(b, 6, 7, 8, 9, 10)
	// remove from a slice
	b = append(b[:2], b[4:]...)
	fmt.Println(b)

	// Maps
	// create a map by `var a map[string]int` <- string is the key type, int is the value type
	c := map[string]int{"foo": 1, "bar": 2}
	// add to a map
	c["baz"] = 3
	fmt.Println(c)
}
