package main

import (
	"fmt"
)

func main() {
	var myMap = map[string]uint8{"Adam": 86, "Eve": 89, "Steve": 99}

	for name, age := range myMap {
		fmt.Printf("Name: %v Age: %v\n", name, age)
	}
	var idx uint8 = 1
	for idx < 11 {
		fmt.Println(idx)
		idx = idx + 1
	}
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
