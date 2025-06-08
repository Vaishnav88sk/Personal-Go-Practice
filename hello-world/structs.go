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
	type carr struct {
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
	myCar := carr{}
	myCar.frontWheel.radius = 5

	// Anonymous struct
	myCary := struct {
		brand string
		model string
	}{
		brand: "Toyota",
		model: "Camry",
	}

	// Embedded structs
	type cars struct {
		brand string
		model string
	}

	type truck struct {
		// "car" is embedded, so the definition of a
		// "truck" now also additionally contains all
		// of the fields of the car struct
		cars
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
	emptyy := emptyStruct{}

}

// area has a receiver of (r rect)
// rect is the struct
// r is the placeholder
func (r rect) area() int {
	return r.width * r.height
}
