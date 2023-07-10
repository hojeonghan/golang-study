package main

import (
	"fmt"

	"github.com/hojeonghan/golang-study/structset"
)

func main() {
	address := structset.Address{Street: "123 Seodaemunro", City: "Seoul", State: "KR", PostalCode: "01234"}
	subscriber := structset.Subscriber{Name: "Aman Singh"}
	subscriber.HomeAddress = address
	fmt.Println(subscriber.HomeAddress)
}
