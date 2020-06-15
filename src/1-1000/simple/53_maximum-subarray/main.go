package main

import (
	"fmt"
)

//func maxSubArray(nums []int) int {
//	var sum, max int
//	for i:=0;i<len(nums) ;i++  {
//		sum = 0
//		for j :=i;j<len(nums);j++{
//			sum +=nums[j]
//			if sum > max{
//				max = sum
//			}
//		}
//	}
//	return max
//}
// sum 只负责相加，当 sum 为负数时，说明子序列中没有发生增益状态，重置起始位置，转向下一个数继续相加。
// max 记录目前子序列相加最大的值为多少。当 sum 超过 max，max 就需要替换为sum的值
func maxSubArray(nums []int) int{
	sum :=0
	max:=nums[0]
	for _, num :=range nums{
		if sum > 0{
			sum += num
		}else {
			sum = num
		}
		if sum > max{
			max = sum
		}
	}
	return max
}
func main() {
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	result := maxSubArray(nums)
	fmt.Println(result)
}
