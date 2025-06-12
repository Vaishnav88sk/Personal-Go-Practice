package main

import "fmt"

func main() {
	// for loop syntax
	/*
		for INITIAL; CONDITION; AFTER{
			// logic
		}
	*/

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// Omitting condition from a for loop
	// for i := 0; ; i++ {
	// 	fmt.Println(i)
	// 	if i == 5{
	// 		return
	// 	}
	// }

	// There is no WHILE loop in Go - same logic done by modifying for loop
	i := 0
	for i < 5 { // for CONDITION{}
		fmt.Println(i)
		i++
	}
	// % - Modulo operator gives remainder after divison .. e.g :- 10%4 = 2
	// && - logical AND operator
	// || - logical OR operator

	// continue keyword - stops thr current iteration and continues to next iteration
	// break keyword - stops thr current iteration and exits

	// for with range
	nums := []int{10, 20, 30}
	for index, value := range nums {
		fmt.Println(index, value)
	}

	// Loop over string
	for i, ch := range "hello" {
		fmt.Println(i, string(ch))
	}

	// Loop over map
	m := map[string]int{"a": 1, "b": 2}
	for key, val := range m {
		fmt.Println(key, val)
	}

}
