package main

/*
1. 플레이어가 추측할 1에서 100사이의 난수를 목표값으로 지정.
2. 플레이어가 추측한 숫자를 입력할 수 있도록 프롬프트를 띄우고 입력받은 추측 값을 저장
3. 플레이어가 추측한 숫자가 목표값보다 낮으면 "Oops, Your guess was LOW."를, 높으면 "Oops, Your guess was HIGH."를 출력
4. 플레이어는 최대 10번까지 추축할 수 있고, 추측할 때마다 남은 횟수를 알려줘야 함
5. 플레이어가 추측한 숫자와 목표값이 같으면 "Good job! You guessed it!" 을 출력하고 프롬프트를 종료
6. 최대 추측 회수 안에 목표값을 맞추지 못하면 "Sorry, You didn't guess my number. It was: [target]" 을 출력
*/
import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func Guess() {
	timeSeed := time.Now().Unix() // 현재 날짜 및 시간을 정수값으로 가져옴.
	rand.Seed(timeSeed)           // timeSeed를 난수의 seed 값으로 사용, 이는 난수 생성 시 무작위성을 높히기 위함.
	target := rand.Intn(100) + 1  // rand는 0~지정한 숫자 사이에 있는 값을 무작위로 출력하는데 실제 1~100사이의 난수가 필요하므로 rand에 +1을 해줌.
	success := false

	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it ? ")
	fmt.Println("당신에겐 총 10번의 기회가 있습니다.")

	for i := 0; i < 10; i++ {
		fmt.Print(i+1, "번 째 추측: ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)  // 입력 받은 문자열 앞 뒤의 공백을 모두 제거.
		guess, err := strconv.Atoi(input) // 입력받은 문자열을 정수로 변환.
		if err != nil {
			log.Fatal(err)
		}

		if target > guess {
			fmt.Println("Oops, Your guess was LOW.")
		} else if target < guess {
			fmt.Println("Oops, Your guess was HIGH.")
		} else {
			success = true
			fmt.Println("Good job! You guessed it!")
			break
		}
	}
	if !success {
		fmt.Println("Sorry, you didn't guess my number. It was: [", target, "]")
	}
}

func main() {
	Guess()
}
