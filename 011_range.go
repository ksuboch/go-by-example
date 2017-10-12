package main

import "fmt"

func main() {

	// range iterates over elements of data structures
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums { // calculating sum of a slice
		sum += num
	}
	fmt.Println(sum)

	for i, num := range nums { // range provides index and value
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs { // range on map iterates over key/value pairs
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs { // range can iterates over keys of a map
		fmt.Println("key:", k)
	}

	for i, c := range "go" { // range on strings iterates over Unicode points
		fmt.Println(i, c)
	}
}
