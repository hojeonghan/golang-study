package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/hojeonghan/golang-study/readfile"
)

// mission#1. 슬라이스를 이용해서 횟수 집계
// func main() {
// 	lines, err := readfile.GetStrings("votes.txt") // 파일에서 문자열을 읽어와서 lines 슬라이스에 저장
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// 후보자 이름과 count를 저장하기 위한 슬라이스 변수 선언
// 	var names []string
// 	var counts []int

// 	matched := false
// 	for _, line := range lines { // liens 슬라이스에서 값을 읽어옴
// 		for i, name := range names {
// 			if name == line { // names 슬라이스에 이미 있는 후보자 이름인 경우
// 				counts[i]++ // count 값을 +1 증가
// 				matched = true
// 			}
// 		}
// 		if !matched {
// 			names = append(names, line) // 새로운 이름인 경우 names 슬라이스에 후보자 이름 추가
// 			counts = append(counts, 1)  // 새로운 이름인 경우 counts 슬라이스 값을 1로 셋팅
// 		}
// 	}
// 	for i, name := range names { // 이름 별 투표수를 출력
// 		fmt.Printf("%s: %d\n", name, counts[i])
// 	}
// }

// mission#2. map을 이용해서 횟수 집계
// func main() {
// 	lines, err := readfile.GetStrings("votes.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	counts := make(map[string]int)
// 	for _, line := range lines {
// 		counts[line]++
// 	}

// 	for key, value := range counts {
// 		fmt.Printf("%s : %d \n", key, value)
// 	}
// }

// mission#3. map을 사용하되 결과가 알파벳 순서로 출력되도록 수정
func main() {
	lines, err := readfile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}

	var names []string
	for name := range counts {
		names = append(names, name)
	}

	sort.Strings(names)
	for _, value := range names {
		fmt.Printf("%s : %d \n", value, counts[value])
	}
}
