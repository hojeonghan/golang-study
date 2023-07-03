package main

//60보다 작은 성적값을 입력하면 failing, 큰 값을 입력하면 passing 메세지를 출력.
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func PassFail() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n') // os.stdin의 값을 \n 전까지 입력받아 input에 저장.
	if err != nil {
		log.Fatal(err)
	}

	// strings.TrimSpace : 문자열의 처음과 끝에 존재하는 모든 공백 문자를 제거.
	input = strings.TrimSpace(input)
	// strconv.ParseFloat : 문자열을 float64 타입으로 변환
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade > 60 {
		status = "passing!"
	} else {
		status = "failing!"
	}
	fmt.Println("A grade of", grade, "is", status)
}

func main() {
	PassFail()
}
