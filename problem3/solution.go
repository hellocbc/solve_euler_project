package main

import "fmt"
import "github.com/hellocbc/solve_euler_project/chocomath"

func main() {
	var targetNumber int64 = 600851475143
	var maxPrime int64 = 0

	for chocomath.IsPrime(targetNumber) != true {
		for i := int64(1); i <= targetNumber; i++ {
			if chocomath.IsPrime(i) == true {
				if targetNumber%i == 0 {
					if i > maxPrime {
						maxPrime = i
					}
					targetNumber /= i
					break
				}
			}
		}
	}
	if targetNumber > maxPrime {
		maxPrime = targetNumber
	}
	fmt.Println(maxPrime)
}
