package main

import (
	"fmt"
)

type Employee struct {
	Name   string
	Salary float64
	Address
}

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
	Address
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}

func main() {
	address := Address{
		Street:     "123 Oak St",
		City:       "Omaha",
		State:      "NE",
		PostalCode: "63111",
	}
	subscriber := Subscriber{
		Name:    "Aman Singh",
		Rate:    8.9,
		Active:  true,
		Address: address,
	}
	fmt.Println(subscriber)
}
