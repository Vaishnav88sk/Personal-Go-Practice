package main

import "fmt"

func main() {
	answer := sub(10, 6)
	fmt.Println(answer)

	// ignore y value
	x, _ := getPoint()
	fmt.Println(x)

	// A return statement without arguments returns the named return values. This is known as a "naked" return.
	a, b := getCoords()
	fmt.Println(a, b)

	// Functions as an arguments
	fmt.Println(aggregate(2, 3, 4, add)) // Function as a value

	// Anonymous function
	newX, newY, newZ := conversions(func(a int) int {
		return a + a
	}, 1, 2, 3)
	// newX is 2, newY is 4, newZ is 6
	fmt.Println(newX, newY, newZ)

	// defer keyword
	// It allows a function to be executed automatically just before its enclosing function returns.
	defer fmt.Println("This is the last statement of program")

	// Go is block scoped - variables scope is inside block {}
}

func sub(x int, y int) int { // Also can be return as (x,y int)
	// int - return type
	return x - y
}

func getPoint() (x int, y int) {
	return 3, 4
}

func getCoords() (x, y int) {
	// x and y are initialized with zero values

	return // automatically returns x and y
}

func aggregate(a, b, c int, arithmetic func(int, int) int) int {
	firstResult := arithmetic(a, b)
	secondResult := arithmetic(firstResult, c)
	return secondResult
}

func add(x, y int) int {
	return x + y
}

func mul(x, y int) int {
	return x * y
}

func conversions(converter func(int) int, x, y, z int) (int, int, int) {
	convertedX := converter(x)
	convertedY := converter(y)
	convertedZ := converter(z)
	return convertedX, convertedY, convertedZ
}
