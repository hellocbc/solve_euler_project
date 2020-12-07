package chocomath

// GetMultiples input nubmers and return output (multiples of input numbers)
func GetMultiples(inputs []int64, maxNumber int64) (output []int64) {
	if len(inputs) == 0 {
		return
	}
	for number := int64(1); number < maxNumber; number++ {
		for _, in := range inputs {
			if number < in {
				continue
			} else if number%in == 0 {
				output = append(output, number)
				break
			}
		}
	}
	return
}

// IsPrime return boolean value, true: prime number, false: non-prime number
func IsPrime(n int64) bool {
	if n == int64(1) {
		return false
	} else if n == int64(2) {
		return true
	}

	halfNum := int64(n / 2)
	for i := int64(2); i <= halfNum; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
