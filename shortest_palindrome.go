package main

import (
	"fmt"
)

func main() {
	var str string = "aaaccaaaa"
	length := len(str)
	result := ""
	i := length - 1
	for ; i > 0; i-- {
		if v1(str, 0, i) {
			if i+1 == length {
				result = str
			} else {
				result = reverse(str[i+1:length]) + str
			}
			break
		}
	}
	if result == "" {
		result = reverse(str[1:length]) + str
	}
	fmt.Print(result)
}

func v1(str string, start, end int) bool {
	for {
		if start < end && str[start] == str[end] {
			start++
			end--
		} else {
			break
		}
	}
	return start >= end
}

func reverse(str string) string {
	length := len(str)
	res := ""
	for i := 0; i < length; i++ {
		res += str[length-i-1 : length-i]
	}
	return res
}
