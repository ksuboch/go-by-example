package main

import "fmt"

func plus(a int, b int) int { // function takes two int and return an int
	return a + b
}

func plusPlus(a, b, c int) int { // function takes three int and return one
	return a + b + c
}

func main() {

	res := plus(1, 2)
	fmt.Println("1 + 2 = ", res) // call a function

	res = plusPlus(1, 2, 3)
	fmt.Println("1 + 2 + 3 = ", res)
}
