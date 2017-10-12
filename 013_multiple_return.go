package main

import "fmt"

func vals() (int, int) { // (int, int) function signature
	return 3, 7 // shows that function returns 2 ints
}

func main() {

	a, b := vals() // 2 different values
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals() // using the blank _ identifier
	fmt.Println(c)
}
