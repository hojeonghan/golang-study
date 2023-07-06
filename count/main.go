package main

import (
	"fmt"
	"log"

	"github.com/hojeonghan/golang-study/readfile"
)

func main() {
	lines, err := readfile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(lines)
}
