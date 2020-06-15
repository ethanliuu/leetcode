package main

import "fmt"

//双指针 替换法
func removeElement(nums []int, val int) int {
	fast := 0
	slow := 0
	for fast < len(nums) {
		// 快指针与 val 值相同时后移。继续寻找一个不与 val 相同的值
		if nums[fast] == val {
			fast++
			// 找到时，替换快慢指针的数值。并且双指针都后移。
		} else if nums[fast] != val {
			nums[slow] = nums[fast]
			fast++
			slow++

		}
	}
	// 返回数组总长度 - 快慢指针位置的差
	return len(nums) - (fast - slow)
}

// 切片法
func removeElement1(nums []int, val int) int {
	for i := 0; i < len(nums); {
		if nums[i] == val {
			// 移除指针当前位置的数值并添加指针之后的数组。
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			// 没问题指针后移
			i++
		}
	}
	return len(nums)
}

func main() {
	nums := []int{3, 2, 2, 3}
	nums1 := []int{3, 2, 2, 3}
	result := removeElement(nums, 3)
	result1 := removeElement1(nums1, 3)
	fmt.Println(result, result1)
}
