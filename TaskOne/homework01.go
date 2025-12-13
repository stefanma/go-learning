package main

import "fmt"

func plusOne(digits []int) []int {
	n := len(digits)

	// 从最后一位开始加1
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	// 如果所有位都变成了0，说明有进位，新增一位1
	return append([]int{1}, digits...)
}

func main() {
	fmt.Println("Hello, World!")
	digits := []int{9, 9, 9}
	result := plusOne(digits)
	fmt.Println("plusOne result is:-----")
	fmt.Println(result)
}
