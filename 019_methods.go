package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int { // this area method has a receiver type of *rect
	return r.width * r.height
}

func (r rect) perim() int { // example of a value receiver (copy on call)
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	// go automatically handles conversion
	// between values and pointers
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())
}
