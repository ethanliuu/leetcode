package main

import "fmt"

func searchInsert(nums []int, target int) int {
	length := len(nums)
	for i:=0;i < length ; i++ {
		// 当 target 与指针所在数值相等时返回下标。当指针所在位置的数值大于 target，即为 target 该插入的位置，也返回下标
		if nums[i] == target||nums[i] > target{
			return i
		}
	}
	// 遍历完整个数组都没有结果。说明target 大于数组所有数值，应该在最后额外插入。
	return length
}
func main() {
	nums:=[]int{1,3,5,6}
	result := searchInsert(nums, 2)
	fmt.Println(result)
}
