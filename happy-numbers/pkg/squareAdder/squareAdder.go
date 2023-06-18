package squareAdder

import (
	"fmt"
	"strconv"
	"sync"
)

func giveDigits(x int) (y []int) {
	xStr := strconv.Itoa(x)
	for _, char := range xStr {
		y = append(y, int(char-'0'))
	}

	return
}

func squareDigits(y []int) (z int) {
	z = 0

	for _, digit := range y {
		z += digit * digit
	}

	return
}

func squareAdder(original int, x int, iteration int) (int, error) {
	xDigits := giveDigits(x)
	result := squareDigits(xDigits)

	if iteration <= 100 {
		if result == 1 {
			return original, nil
		} else if result != x {
			iteration++
			squareAdder(original, result, iteration)
		}
	}
	return 0, fmt.Errorf("number %d has had 100 iterations and wasn't able to arrive at 1", x)
}

func InvokeChecker(id int, inputChn chan int, resultChn chan int, wg *sync.WaitGroup) {
	for number := range inputChn {
		// fmt.Printf("worker: %d, number: %d\n", id, number)
		result, err := squareAdder(number, number, 0)
		if err == nil {
			resultChn <- result
		}
	}
	wg.Done()
}
