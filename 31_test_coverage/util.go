package util

import (
	"errors"
)

// ReverseInt - function to reverse int. For example 123 -> 321
func ReverseInt(x interface{}) (int, error) {
	y, ok := x.(int)
	if !ok {
		return 0, errors.New("not int")
	}

	var result int
	for y != 0 {
		result = result*10 + y%10
		if result > 2147483647 || result < -2147483648 {
			return 0, nil
		}

		y /= 10
	}

	return result, nil
}

// ContainsDuplicate - check int array for duplicate. For example [1, 1, 2, 3] -> true
func ContainsDuplicate(nums []int) bool {
	set := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := set[v]; ok {
			return true
		}

		set[v] = struct{}{}
	}

	return false
}

// IsPalindrome - check integer for polindrome. For example 101 -> true, 123 -> false
func IsPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	var revertedNum int
	for x > revertedNum {
		last := x % 10
		x /= 10
		revertedNum = revertedNum*10 + last
	}

	return x == revertedNum || x == revertedNum/10
}

// Fib - calculate fibanachi num. For example 10 -> 34
func Fib(n int) int {
	if n < 2 {
		return n
	}

	return Fib(n-1) + Fib(n-2)
}

// MakeSlice - make new slice with size and put zero for all cell. For example 3 -> [0, 0, 0]
func MakeSlice(l int) []int {
	a := make([]int, 0)
	for i := 0; i < l; i++ {
		a = append(a, i)
	}

	return a
}

// Pad - make new string with template data. For example ("a", 3, "b") -> "abbb"
func Pad(s string, length int, template string) string {
	for len(s) < length {
		s += template
	}

	return s
}
