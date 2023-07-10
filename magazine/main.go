package main

import (
	"fmt"

	"github.com/hojeonghan/golang-study/structset"
)

func main() {
	subscriber := structset.Subscriber{Name: "Aman Singh"}
	subscriber.Address.Street = "123 Oak St"
	subscriber.Address.City = "Omaha"
	subscriber.State = "NE"
	subscriber.PostalCode = "63111"
	fmt.Println("Street: ", subscriber.Street)
	fmt.Println("Street: ", subscriber.Street)
}
