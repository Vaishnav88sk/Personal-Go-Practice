package main

import "fmt"

func main() {
	// if statements in Go do not use parentheses
	age := 10
	if age > 18 {
		fmt.Println("You are eligible to vote.")
	} else if age <= 14 {
		fmt.Println("You are teenager.")
	} else {
		fmt.Println("You are not eligible to vote.")
	}

	// An if conditional can have an "initial" statement
	if length := 10; length < 12 {
		fmt.Println("Email is invalid")
	}

	// Switch case
	/*
		The break statement is not required at the end of a case.
		The break statement is implicit in Go.
		If you do want a case to fall through to the next case, you can use the 'fallthrough' keyword.
	*/

	owner := getCreator("macOS")
	fmt.Println("Creator is:", owner)

}
func getCreator(os string) string {
	var creator string
	switch os {
	case "linux":
		creator = "Linus Torvalds"
	case "windows":
		creator = "Bill Gates"

	// all three of these cases will set creator to "A Steve"
	case "macOS":
		fallthrough
	case "Mac OS X":
		fallthrough
	case "mac":
		creator = "A Steve"

	default:
		creator = "Unknown"
	}
	return creator
}
