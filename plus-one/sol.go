package main

func plusOne(digits []int) []int {
	carry := 0
	for i := len(digits) - 1; 0 <= i; i-- {
		sum := digits[i] + carry
		if i == len(digits)-1 {
			sum++
		}
		if sum <= 9 {
			digits[i] = sum
			return digits
		} else {
			digits[i] = sum % 10
			carry = (sum - (sum % 10)) / 10
		}
	}
	return append([]int{carry}, digits...)
}
