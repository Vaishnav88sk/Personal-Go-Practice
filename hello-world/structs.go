package main

import "fmt"

func main() {
	// Structs - group different types of variables together
	type car struct {
		brand   string
		model   string
		doors   int
		mileage int
	}

	// Nested structs
	type car struct {
		brand      string
		model      string
		doors      int
		mileage    int
		frontWheel wheel
		backWheel  wheel
	}

	type wheel struct {
		radius   int
		material string
	}
	myCar := car{}
	myCar.frontWheel.radius = 5

	// Anonymous struct
	myCar := struct {
		brand string
		model string
	}{
		brand: "Toyota",
		model: "Camry",
	}

	// Embedded structs
	type car struct {
		brand string
		model string
	}

	type truck struct {
		// "car" is embedded, so the definition of a
		// "truck" now also additionally contains all
		// of the fields of the car struct
		car
		bedSize int
	}

	// Structs menthods
	type rect struct {
		width  int
		height int
	}

	// anonymous empty struct type
	empty := struct{}{}

	// named empty struct type
	type emptyStruct struct{}
	empty := emptyStruct{}

}

// area has a receiver of (r rect)
// rect is the struct
// r is the placeholder
func (r rect) area() int {
	return r.width * r.height
}
