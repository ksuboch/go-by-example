package main

import "fmt"

func main() {

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ { // initial/condition/after for loop
		fmt.Println(j)
	}

	for { // repeat until the break
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue // continue to the next iteration of the loop
		}
		fmt.Println(n)
	}
}
