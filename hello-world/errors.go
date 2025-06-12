package main

import (
	"errors"
	"fmt"
)

func main() {
	// Error handling - the error Interface

	_, err1 := getUser()
	if err1 != nil {
		fmt.Println("Error occured!!")
		return
	}
	// Go doesn’t have exceptions. Instead, functions return errors as their last return value.

	result, err2 := divide(10, 5)
	if err2 != nil {
		fmt.Println("Error:", err2)
		return
	}
	fmt.Println("Result:", result)

	// Can build custom types that implements the error interface

	// New() function for defining new error
	var err3 error = errors.New("something went wrong!")
	fmt.Println(err3)

	// 'errors.Is' and 'errors.As'
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File does not exist")
	}

	var pathErr *os.PathError
	if errors.As(err, &pathErr) {
		fmt.Println("It's a path error")
	}

}

func getUser() (string, error) {
	return "", nil
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// panic and recover  - same like try and catch
func risky() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	panic("something went wrong")
}

/*
Use panic only in unrecoverable situations (like nil pointer dereference, bad system state).
Can use recover() in defer to catch panics, but it's rare and discouraged in most normal Go programs.
*/
