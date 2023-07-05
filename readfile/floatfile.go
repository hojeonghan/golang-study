// Package readfile은 파일로부터 샘플 데이터를 읽어 옵니다.
package readfile

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

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

func SliceWithFile(filename string) ([]float64, error) {
	var numbers []float64
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		numbers = append(numbers, number)
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
