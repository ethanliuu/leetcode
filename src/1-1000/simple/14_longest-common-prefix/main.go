package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs)==0{
		return ""
	}
	// 首先拿 strs[0] 作为基准。
	pre := strs[0]
	for i:=1;i<len(strs) ;i++  {
		// 截取下一个元素。
		cur := strs[i]
		// 判断基准元素在不在下一个元素中
		for strings.Index(cur,pre)!=0 {
			// 如果将基准元素切完还没找到共同的子串，返回""。
			if len(pre)==0{
				return ""
			}
			// 将基准元素最后的字符切掉，保留前面字符的进行下次比较。
			pre = pre[:len(pre)-1]
		}
	}
	return pre
}
func main() {
	strs := []string{"flower","flow","flight"}
	result:= longestCommonPrefix(strs)
	fmt.Println(result)
}
