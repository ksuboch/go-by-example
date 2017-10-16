package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20}) // new struct

	fmt.Println(person{name: "Alice", age: 30}) // name field when initializing

	fmt.Println(person{name: "Fred"}) // omitted field will be zero-valued

	fmt.Println(&person{name: "Ann", age: 40}) // an & prefix yeald pointer to the struct

	s := person{name: "Sean", age: 50} // access field with a dot
	fmt.Println(s.name)

	sp := &s            // using dots with struct pointers
	fmt.Println(sp.age) // pointer will be dereferenced

	sp.age = 51 // structs are mutable
	fmt.Println(sp.age)
}
