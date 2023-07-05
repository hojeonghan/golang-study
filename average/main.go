/*
- 3주 동안 사용한 소고기 양에 대한 샘플 데이터를 기반으로 이번 주에 주문할 소고기 양을 계산하는 프로그램.
- 샘플 데이터의 평균을 구해 이번주의 적절한 추정치를 구한다.
*/
package main

import (
	"fmt"
	"github.com/golang-study/readfile"
	"log"
)

func main() {
	// 이미 3주간 사용한 소고기 양을 알고 있으므로 배열 리터럴을 사용.
	// weights := [3]float64{71.8, 56.2, 89.5}
	// // 배열의 타입이 float64이므로 나중에 평균값을 구하기 위해 sampleCount 수도 float64로 형변환.
	// sampleCount := float64(len(weights))
	// var sum float64 = 0
	// for _, weight := range weights {
	// 	sum += weight
	// }

	// fmt.Printf("Average is: %0.2f\n", sum/sampleCount)
	numbers, err := readfile.ReadFile("sampledata.txt")
	if err != nil {
		log.Fatal(err)
	}

	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}

	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)
}
