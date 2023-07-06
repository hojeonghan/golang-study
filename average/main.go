/*
- 3주 동안 사용한 소고기 양에 대한 샘플 데이터를 기반으로 이번 주에 주문할 소고기 양을 계산하는 프로그램.
- 샘플 데이터의 평균을 구해 이번주의 적절한 추정치를 구한다.
*/
package main

import (
	"fmt"
	"log"

	"github.com/hojeonghan/golang-study/readfile"
)

// 파일에서 읽어온 데이터로 array 만들어 사용하는 예제
func main() {
	numbers, err := readfile.SliceWithFile("sampledata.txt")
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
