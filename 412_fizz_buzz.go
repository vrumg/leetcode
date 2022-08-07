package leetcode

import (
	"strconv"
)

func fizzBuzz(n int) []string {
	result := make([]string, n)

	for i := 1; i <= n; i++ {

		sum := sumOfDigits(i)
		lastDigit := i % 10
		if sum%3 == 0 {
			if lastDigit == 0 || lastDigit == 5 {
				result[i-1] = "FizzBuzz"
				continue
			} else {
				result[i-1] = "Fizz"
				continue
			}
		}
		if lastDigit == 0 || lastDigit == 5 {
			result[i-1] = "Buzz"
			continue
		}
		result[i-1] = strconv.Itoa(i)
	}

	return result
}

func sumOfDigits(n int) int {
	var count int
	var remainder int

	for n != 0 {
		remainder = n % 10
		count += remainder
		n = n / 10
	}

	return count
}
