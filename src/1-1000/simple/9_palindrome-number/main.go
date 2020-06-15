package main

import (
	"fmt"
)
// 将 x 倒序成 y，判断与正序是否一样。
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	y := 0
	re := x
	for re != 0 {
		y = y*10 + re%10
		re /= 10
	}
	return x == y
}

func main() {
	result := isPalindrome(123321)
	fmt.Println(result)
}
