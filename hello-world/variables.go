package main

import "fmt"

func main() {
	// This is printing line
	fmt.Println("Hello World!")

	// This is for formatted print line
	fmt.Printf("hello World!\n")

	// This is varibale declaration
	var a int = 10
	var valid bool = true
	var word string = "I'm learning Go..."
	var price float64 = 25.0
	var b byte = 10
	var def int

	fmt.Printf("This is int %v\n", a)
	fmt.Printf("This is bool %v\n", valid)
	fmt.Printf("This is string %q\n", word)
	fmt.Printf("This is price %.2f\n", price)
	fmt.Printf("This is byte %v\n", b)
	fmt.Printf("This is default value: %v of variable def\n", def)
	fmt.Printf("This is combination: %v %q\n", a, word)

	// GOATed variable declaration using walrus operator (:=) - for dynamic variable
	message := "Hello guys"
	msg := 24

	fmt.Printf("Dynamic variables: %v %v\n", message, msg)
	fmt.Println("Dynamic varibales:", message, msg)
	fmt.Println("Dynamic varibales:", message+": "+word) // + for same type

	// variables types can be depending on sizes - int64, int32, uint32 (unsigned), complex128 (complex number having real and imaginary part)

	// declare multiple variables on the same line
	msd, vk := 7, 18
	fmt.Println("MSD:", msd, "VK:", vk)

	// Constants are declared with the const keyword. They can't use the := short declaration syntax
	const pi = 3.14159
	fmt.Println("Constant:", pi)

	// you cannot declare a constant that can only be computed at run-time like you can in JavaScript

	// fmt.Sprintf() - Returns the formatted string
	s := fmt.Sprintf("The value is %v", 10)
	fmt.Println(s)
	/*
		String - %s
		Integer - %d
		Float - %f , %.2f
	*/
	const name = "Saul Goodman"
	const openRate = 30.5
	mseg := fmt.Sprintf("Hi %s, your open rate is %.1f percent\n", name, openRate)
	fmt.Print(mseg)

	/*
		Go also has a special type, rune, which is an alias for int32.
		This means that a rune is a 32-bit integer, which is large enough to hold any Unicode code point.
	*/
}
