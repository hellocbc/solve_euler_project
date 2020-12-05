package main

import "fmt"

func main() {
	var inputNumber = []int64{3, 5}
	var maxNumber int64 = 1000

	var sumNumber int64 = 0
	for i := int64(1); i < maxNumber; i++ {
		for _, n := range inputNumber {
			if i%n == 0 {
				sumNumber += i
				break
			}
		}
	}
	fmt.Println(sumNumber)
}
