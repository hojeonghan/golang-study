// Package structset magazine 구현에 공용으로 사용되는 구조체를 선언합니다.
package structset

type Employee struct {
	Name        string
	Salary      float64
	HomeAddress Address
}

type Subscriber struct {
	Name        string
	Rate        float64
	Active      bool
	HomeAddress Address
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}
