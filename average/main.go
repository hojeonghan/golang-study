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

// #1.mission) 파일에서 읽어온 데이터로 array 만들어 사용하는 예제
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

// #2.mission) 명령줄 인자로 데이터를 받도록 수정
// func main() {
// 	dataList := os.Args[1:]
// 	fmt.Println(dataList)

// 	var sum float64 = 0
// 	for _, element := range dataList {
// 		number, err := strconv.ParseFloat(element, 64)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		sum += number
// 	}

// 	dataCount := float64(len(dataList))
// 	fmt.Printf("Average: %0.2f\n", sum/dataCount)
// }

// #3. mission) 가변 인자 함수로 데이터를 받아 평균을 구하기
// func average(datalist ...float64) float64 {
// 	var sum float64 = 0
// 	for _, data := range datalist {
// 		sum += data
// 	}

// 	return sum / float64((len(datalist)))
// }
// func main() {
// 	dataList := os.Args[1:]
// 	var numbers []float64
// 	for _, element := range dataList {
// 		number, err := strconv.ParseFloat(element, 64)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		numbers = append(numbers, number)
// 	}
// 	fmt.Printf("Average: %0.2f\n", average(numbers...))
// }
