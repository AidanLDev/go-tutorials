package main

import "fmt"

func main() {
	var p *int32 = new(int32)
	var i int32

	p = &i
	*p = 1

	fmt.Printf("The value of p points to %v", *p)
	fmt.Printf("\nThe value of i is: %v", i)
}
