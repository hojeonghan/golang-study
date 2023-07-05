package main

import (
	"fmt"
)

func Array() {
	// 이미 3주간 사용한 소고기 양을 알고 있으므로 배열 리터럴을 사용.
	weights := [3]float64{71.8, 56.2, 89.5}
	// 배열의 타입이 float64이므로 나중에 평균값을 구하기 위해 sampleCount 수도 float64로 형변환.
	sampleCount := float64(len(weights))
	var sum float64 = 0
	for _, weight := range weights {
		sum += weight
	}

	fmt.Printf("Average is: %0.2f\n", sum/sampleCount)
}
