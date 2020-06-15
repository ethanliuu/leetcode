package main

import "fmt"

func match(s string) string {
	var res string
	switch s {
	case ")":
		res = "("
	case "}":
		res = "{"
	case "]":
		res = "["
	}
	return res
}

func isValid(s string) bool {
	stack := make([]string, 0)
	// 遍历字符串
	for _, i := range s {
		switch string(i) {
		// 如果是左括号，将元素添加到待处理队列中。
		case "(", "{", "[":
			stack = append(stack, string(i))
			// 如果是右括号
		case ")", "}", "]":
			// 先判断 stack 里有没有待处理元素，然后判断当前右括号的映射与 stack 末元素是否相同。
			if len(stack) != 0 && stack[len(stack)-1] == match(string(i)) {
				// 符合之后，切掉待处理元素队列中的最后一个元素。
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func main() {
	s := "({[]})"
	result := isValid(s)
	fmt.Println(result)
}
