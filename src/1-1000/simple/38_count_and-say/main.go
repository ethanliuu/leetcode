package main

import (
	"bytes"
	"fmt"
	"strconv"
)

//1.     1
//2.     11
//3.     21
//4.     1211
//5.     111221
//6.	 312211
//7.     13112221
//8.     1113213211
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	// 递归，获取上次是结果
	pre := countAndSay(n - 1)
	// 定义计数器，统计相同数字出现次数
	count := 1
	// 当前要判断的数字
	var cur = pre[0]
	var bt bytes.Buffer
	for i := 0; i < len(pre)-1; i++ {
		// 如果当前数字和下一个数字相同，计数器++
		if pre[i] == pre[i+1] {
			count++
		} else {
			// 如果不相同，将计数器的值和上个记录的值写入buffer中。
			bt.WriteString(strconv.Itoa(count))
			bt.WriteString(string(pre[i]))
			cur = pre[i+1]
			count = 1
		}
	}
	// 最后整合，写入到记录中
	bt.WriteString(strconv.Itoa(count))
	bt.WriteString(string(cur))
	return bt.String()
}

func main() {
	result := countAndSay(5)
	fmt.Println(result)
}
