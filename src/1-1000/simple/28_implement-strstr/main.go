package main

import "fmt"

func strStr(haystack string, needle string) int {
	l1 := len(haystack)
	l2 := len(needle)
	if l2 == 0 {
		return 0
	}
	if l2 > l1 || l1 == 0 {
		return -1
	}
	for i := 0; i <= l1-l2; i++ {
		// 截取指针当前位置加上 l2 长度的切片是否是 l2
		if haystack[i:i+l2] == needle {
			return i
		}
	}
	return -1
}
func main() {
	result := strStr("hello", "ll")
	fmt.Println(result)
}
