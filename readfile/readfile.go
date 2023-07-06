// Package readfile은 파일로부터 샘플 데이터를 읽어 옵니다.
package readfile

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// 3개의 float64 타입의 데이터를 가진 파일을 인자로 받습니다.
// 파일로부터 읽은 데이터를 배열로 반환합니다.
func ReadFloatFile(filename string) ([3]float64, error) {
	var numbers [3]float64
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		i++
	}
	err = file.Close()
	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}

	return numbers, nil
}

// n개의 float64 타입의 데이터를 가진 파일을 인자로 받습니다.
// 파일로부터 읽은 데이터를 슬라이스로 반환합니다.
func SliceWithFile(filename string) ([]float64, error) {
	var numbers []float64
	file, err := os.Open(filename)
	if err != nil {
		return nil, err // 슬라이스가 유효하지 않은 경우 nil을 반환
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err // 슬라이스가 유효하지 않은 경우 nil을 반환
		}
		numbers = append(numbers, number)
	}
	err = file.Close()
	if err != nil {
		return nil, err // 슬라이스가 유효하지 않은 경우 nil을 반환
	}
	if scanner.Err() != nil {
		return nil, scanner.Err() // 슬라이스가 유효하지 않은 경우 nil을 반환
	}

	return numbers, nil
}
