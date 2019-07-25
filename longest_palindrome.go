package main

import (
	"fmt"
)

func main() {
	fmt.Println(version_2())
}

//从中间往两边发散
func version_2() string {
	str := "babdd"

	length := len(str)
	start, end := 0, 0
	for j := 0; j < length; j++ {
		funcv2(str, length, j, j, &start, &end)
		funcv2(str, length, j, j+1, &start, &end)
	}
	return str[start:end]

}
func funcv2(s string, length, i, j int, start, end *int) {
	for {
		if i >= 0 && j < length && s[i] == s[j] {
			i--
			j++
		} else {
			break
		}
	}
	if *end - *start < j-i-1 {
		*start = i + 1
		*end = j
	}
	return
}

//往中间集合
func version_1() string {
	str := "adcbbbidesd"

	length := len(str)

	ret := ""
	var k int
	tempRet := ""
	for j := 0; j < length; j++ {
		newround := true
		j1 := j
		endIndex := length
		for k = length - 1; k >= j; k-- {
			if k < j1 {
				break
			}
			if str[k] == str[j1] {
				if newround {
					endIndex = k
					tempRet = str[j1 : k+1]
					newround = false
				}
				j1++
			} else {
				if tempRet != "" {
					k = endIndex
					endIndex = length
				}
				tempRet = ""
				newround = true
				j1 = j
			}
		}
		if len(tempRet) > len(ret) {
			ret = tempRet
		}
		k = length - 1
	}
	return ret
}
