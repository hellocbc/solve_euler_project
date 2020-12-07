package chocomath_test

import (
	"reflect"
	"testing"

	"github.com/hellocbc/solve_euler_project/chocomath"
)

type TestCaseMultiples struct {
	nums         []int64
	maxNum       int64
	expectedNums []int64
}

func TestMultipels(t *testing.T) {
	var testCases []TestCaseMultiples
	testCases = append(testCases, TestCaseMultiples{[]int64{3, 5}, 10, []int64{3, 5, 6, 9}})
	testCases = append(testCases, TestCaseMultiples{[]int64{3, 5, 7}, 20, []int64{3, 5, 6, 7, 9, 10, 12, 14, 15, 18}})

	for _, in := range testCases {
		if reflect.DeepEqual(chocomath.GetMultiples(in.nums, in.maxNum), in.expectedNums) == false {
			t.Errorf("Wrong Result: %d vs %d", chocomath.GetMultiples(in.nums, in.maxNum), in.expectedNums)
		}
	}
}

type TestCasePrime struct {
	prime   int64
	isPrime bool
}

func TestIsPrime(t *testing.T) {
	var testCases []TestCasePrime
	testCases = append(testCases, TestCasePrime{1, false})
	testCases = append(testCases, TestCasePrime{2, true})
	testCases = append(testCases, TestCasePrime{3, true})
	testCases = append(testCases, TestCasePrime{4, false})
	testCases = append(testCases, TestCasePrime{11, true})
	testCases = append(testCases, TestCasePrime{113, true})
	testCases = append(testCases, TestCasePrime{222, false})
	testCases = append(testCases, TestCasePrime{599, true})
	testCases = append(testCases, TestCasePrime{771, false})
	testCases = append(testCases, TestCasePrime{991, true})

	for _, in := range testCases {
		if chocomath.IsPrime(in.prime) != in.isPrime {
			t.Errorf("Wrong Result: %d -> %t vs %t", in.prime, chocomath.IsPrime(in.prime), in.isPrime)
		}
	}
}
