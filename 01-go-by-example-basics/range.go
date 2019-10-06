package main

import "fmt"

func main() {

	// range on arrays and slices provides both the "index" and "value" for each entry
	nums := []int{2, 3, 4}
	for index, num := range nums {
		if num == 3 {
			fmt.Println("Index:", index)
		}
	}

	// Here we didn't need the "index", so we ignored it with the blank identifier "_"
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("Sum:", sum)

	// range on maps iterates over key/value pairs
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for key, value := range kvs {
		fmt.Printf("%s -> %s\n", key, value)
	}

	// range can also iterate over just the keys of a map
	for key := range kvs {
		fmt.Println("Key:", key)
	}

	// A Unicode code point is a unique number assigned to each Unicode character
	// range on strings iterates over Unicode code points
	// The first value is the starting byte index of the rune and the second the rune itself
	for index, character := range "go" {
		fmt.Println(index, character)
	}
}
