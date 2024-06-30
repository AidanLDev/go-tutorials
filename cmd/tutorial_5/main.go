package main

import "fmt"

type owner struct {
	name string
}

type gasEngine struct {
	mpg     uint8
	gallons uint8
	owner   owner
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it!")
	} else {
		fmt.Println("Fill up first mate!")
	}
}

// Creating a method
func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

func main() {
	var myEngine gasEngine
	fmt.Println(myEngine.mpg, myEngine.gallons)

	var hisEngine gasEngine = gasEngine{mpg: 50, gallons: 75}
	fmt.Println(hisEngine)

	// field names can be omitted and they'll be set in order
	var theirEngine gasEngine = gasEngine{22, 52, owner{"Harry"}}
	fmt.Println(theirEngine)

	// fields can be set after instillation
	myEngine.gallons = 99
	myEngine.mpg = 15
	myEngine.owner.name = "Aidan"
	fmt.Printf("Total miles left in le tank %v", myEngine.milesLeft())
	fmt.Println(myEngine)

	var grettasCar = struct {
		engineType string
		mpg        uint8
		capacity   uint8
	}{"Electric", 100, 55}

	fmt.Println(grettasCar)

	var newEngine electricEngine = electricEngine{25, 15}
	canMakeIt(newEngine, 55)

}
