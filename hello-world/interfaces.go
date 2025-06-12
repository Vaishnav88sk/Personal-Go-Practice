package main

import "fmt"

func main() {
	// An interface in Go is a type that defines a set of method signatures.
	// Any type that implements all those methods implicitly satisfies the interface.
	d := Dog{}
	r := Robot{}
	announce(d) // Woof!
	announce(r) // Beep!
	// Both Dog and Robot implement the Speak() method

	// Interfaces are implemented implicitly - No need to write implements like Java
	// Empty interfaces - with zero methods - Every type satisfies it - equivalent to 'any'

	// type switch - similar to switch satement but cases specify types
	var s interface{} = "hello"

	str, ok := s.(string) // type assertion
	if ok {
		fmt.Println("It's a string:", str)
	}

	switch v := s.(type) { // type switch
	case string:
		fmt.Println("It's a string:", v)
	case int:
		fmt.Println("It's an int:", v)
	}

}

type Speaker interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

type Robot struct{}

func (r Robot) Speak() string {
	return "Beep!"
}

func announce(s Speaker) {
	fmt.Println(s.Speak())
}
