package main

import "fmt"
import "strconv"
import s "strings"

func isPalindrome(n int) bool {
	var strNum string = strconv.Itoa(n)
	lenStr := len(strNum)
	//fmt.Println(lenStr)
	//fmt.Println(s.Split(strNum, ""))

	halfLen := lenStr / 2
	arrStr := s.Split(strNum, "")
	for i := 0; i < halfLen; i++ {
		if arrStr[i] != arrStr[lenStr-1-i] {
			return false
		}
	}

	return true
}

func main() {
	var maxPalindrome int
	var maxI int
	var maxJ int
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			multiNum := i * j
			if isPalindrome(multiNum) == true && multiNum > maxPalindrome {
				maxPalindrome = multiNum
				maxI = i
				maxJ = j
			}
		}
	}
	fmt.Printf("%d * %d = %d\n", maxI, maxJ, maxPalindrome)
}
