package main

import (
	"errors"
	"fmt"
)

func main() {
	printName("Aidan")
	result, remainder, err := intDivision(55, 0)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("The result is %v and the remainder is %v...", result, remainder)
}

func printName(name string) {
	fmt.Println(name)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	return numerator / denominator, numerator % denominator, err
}
