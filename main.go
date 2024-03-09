package main

import "fmt"

func main() {
	result := add(1, 2)
	println(result)

	a, err := divide(10, 0)
	if err != nil {
		println(err.Error())
	} else {
		println(a)
	}
}

// define a function that takes two integers and returns an integer
func add(x, y int) int {
	return x + y
}

// error handling in functions
func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return x / y, nil
}
