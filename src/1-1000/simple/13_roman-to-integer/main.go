package main

import "fmt"
// pre 的加减取决于 cur ，cur 小于 pre 则相加，反之相减。
func romanToInt(s string) int {
	mp := map[string]int{
		"I":1,
		"V":5,
		"X":10,
		"L":50,
		"C":100,
		"D":500,
		"M":1000,
	}
	var sum int
	pre := mp[string(s[0])]
	for i := 1; i < len(s); i++ {
		cur := mp[string(s[i])]
		if pre>=cur{
			sum+=pre
		}else {
			sum-=pre
		}
		pre=cur
	}
	sum += mp[string(s[len(s)-1])]
	return sum
}

func main() {
	result := romanToInt("IX")
	fmt.Println(result)
}
