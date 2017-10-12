package main

import "fmt"

func main() {

	m := make(map[string]int) // to create an empty map use builtin 'make'

	m["k1"] = 7 // setting key/value pairs using name[key] = val syntax
	m["k2"] = 13

	fmt.Println("map:", m) // show all key/value pairs

	v1 := m["k1"] // getting value for a key 'k1'
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m)) // number of key/value pairs

	delete(m, "k2") // removing key/value pair
	fmt.Println("map:", m)

	_, prs := m["k2"] // the optional second return value if the key was in map
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2} // declare and initialize
	fmt.Println("map:", n)
}
