package main

import "fmt"
//采用双指针 替换法
func removeDuplicates(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	//起始位置
	p := 0
	q := 0
	// 遍历数组
	for q != len(nums) {
		// 如果 p 指向的值和 q 指向的值一样，q 后移
		if nums[p] == nums[q] {
			q++
			// 如果 q 指向的值大于 p 指向的值，将 p+1 的指向的值改为 q 的值。
		} else if nums[q] > nums[p] {
			// 只有出现重复的值才修改源数组
			if q-p > 1 {
				nums[p+1] = nums[q]
			}
			p++
		}
	}
	return p + 1

}

// 切片法
func removeDuplicates1(nums []int) int {
	for i := 0; i < len(nums); {
		if nums[i]==nums[i+1]{
			nums = append(nums[:i],nums[i+1:]...)
		}else {
			i++
		}
	}
	return len(nums)
}
func main() {
	nums := []int{1, 2, 2, 3, 4, 4, 5}
	result := removeDuplicates(nums)
	for i:=0;i<result ;i++  {
		fmt.Println(nums[i])
	}
	fmt.Println(nums)

}
