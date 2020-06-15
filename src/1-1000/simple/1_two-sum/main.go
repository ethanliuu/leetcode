package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if k, ok := m[target-v]; ok {
			return []int{k, i}
		}
		m[v] = i
	}
	return nil
}

func main() {
	var nums = []int{2, 7, 11, 15}
	result := twoSum(nums, 9)
	fmt.Println(result)
}