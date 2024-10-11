package main

import "fmt"

type gasEngine struct {
	mpg       uint8
	gallons   uint8
	ownerInfo owner
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

type owner struct {
	name string
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.mpkwh * e.kwh
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

func main() {
	var myEngine gasEngine = gasEngine{mpg: 25, gallons: 12, ownerInfo: owner{name: "Segun"}}
	var my2Engine electricEngine = electricEngine{12, 10}
	myEngine.mpg = 20
	myEngine.ownerInfo.name = "GY"
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerInfo.name)
	fmt.Printf("Total miles left in tank: %v\n", myEngine.milesLeft())

	canMakeIt(myEngine, 50)
	canMakeIt(my2Engine, 200)

}
