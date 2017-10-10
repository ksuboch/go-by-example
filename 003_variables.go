package main

import "fmt"

func main() {

	var a string = "initial"
	fmt.Println(a)

	var b, c int = 1, 2 // declaring multiple variables
	fmt.Println(b, c)

	var d = true // Go will infer the type
	fmt.Println(d)

	var e int // zero-valued variable
	fmt.Println(e)

	f := "short" // declaring and initializing e.g. for var f strion = "short"
	fmt.Println(f)
}
