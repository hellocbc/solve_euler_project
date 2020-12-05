package main

import "fmt"

func main() {
	var firstNumber int64 = 0
	var secondNumber int64 = 1
	var currentValue int64 = 0
	var maxValue int64 = 4000000
	var sumValue int64 = 0

	for currentValue < maxValue {
		if currentValue%2 == 0 {
			//fmt.Println(currentValue)
			sumValue += currentValue
		}
		currentValue = firstNumber + secondNumber
		firstNumber = secondNumber
		secondNumber = currentValue
	}

	fmt.Println(sumValue)
}
