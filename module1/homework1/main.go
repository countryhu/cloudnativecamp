package main

import "fmt"

/*
课后练习 1.1
编写一个小程序：
给定一个字符串数组
[“I”,“am”,“stupid”,“and”,“weak”]
用 for 循环遍历该数组并修改为
[“I”,“am”,“smart”,“and”,“strong”]
*/
func main() {
	strArr := []string{"I", "am", "stupid", "and", "weak"}
	for i := 0; i < len(strArr); i++ {
		if strArr[i] == "stupid" {
			strArr[i] = "smart"
		} else if strArr[i] == "weak" {
			strArr[i] = "strong"
		}
	}
	fmt.Println(strArr)
}
