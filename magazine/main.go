package main

import (
	"fmt"

	"github.com/hojeonghan/golang-study/structset"
)

func main() {
	var employee structset.Employee
	employee.Name = "Joy Carr"
	employee.Salary = 60000
	fmt.Println(employee.Name)
	fmt.Println(employee.Salary)
}
