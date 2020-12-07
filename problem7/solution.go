package main

import "github.com/hellocbc/solve_euler_project/chocomath"

import "fmt"

func main() {
	var numPrimes int
	var baseNumber int = 1

	for numPrimes < 10001 {
		baseNumber++
		if chocomath.IsPrime(int64(baseNumber)) == true {
			numPrimes++
		}
	}
	fmt.Println(baseNumber)
}
