package main

import "fmt"

func main() {
	var intNum int = 3200
	intNum += 1

	var floatNum float64 = 100.5

	fmt.Println(intNum)
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	// Strings
	var myString string = "Hello mates"
	var myDoubleLineString string = `Double
	Liner`
	fmt.Println(myString)
	fmt.Println(len(myDoubleLineString))

	var myRune rune = 'a'
	fmt.Println(myRune)
	myNameIs := "Aidan"
	fmt.Println(myNameIs)

	const myLastName string = "Lowson"
	fmt.Println(myLastName)
}
