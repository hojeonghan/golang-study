package main

import (
	"fmt"
)

type Liters float64
type Gallons float64
type MyType string

type Number int

func ToLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}

func ToGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}

func (m MyType) sayKimchi() {
	fmt.Println(m, ",say Kimchi!!")
}

func (n *Number) Double() {
	*n *= 2
	fmt.Println("output of Double is:", *n)
}

func main() {
	// var carFuel Gallons
	// var busFuel Liters

	// carFuel = Gallons(1.2)
	// busFuel = Liters(4.5)
	// carFuel += ToGallons(Liters(40.0))
	// busFuel += ToLiters(Gallons(30.0))

	// fmt.Printf("Car fuel: %0.1f gallons\n", carFuel)
	// fmt.Printf("Bus fuel: %0.1f liters\n", busFuel)

	// value := MyType("Peter")
	// value.sayKimchi()
	// anotherValue := MyType("Amy")
	// anotherValue.sayKimchi()

	number := Number(4)
	fmt.Println("original value of number:", number)
	number.Double()
	fmt.Println("number after calling Double:", number)
}
