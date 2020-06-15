package main

import "fmt"

func reverse(x int) int {
	y := 0
	for x != 0 {
		y = y*10 + x%10
		// 移位运算 2的31次方
		if !((-1<<31) <= y || y <= (1<<31)-1) {
			return 0
		}
		x /= 10
	}
	return y
}

func main() {
	result := reverse(1534236469)
	fmt.Println(result)
}
