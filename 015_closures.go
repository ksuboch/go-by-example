package main

import "fmt"

func intSeq() func() int { // return another function
	i := 0
	return func() int { // closes the variable i
		i += 1
		return i
	}
}

func main() {

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
