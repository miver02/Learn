package main

import "fmt"

func main() {
	timeStr1, timeStr2 := "2025/10/01 0:09:20", "2025/10/01 0:23:07"
	timeSub, err := timeSub(timeStr1, timeStr2)
	if err != nil || timeSub == -1 {
		fmt.Printf("系统错误: %v\n", err)
		return
	} else {
		fmt.Printf("%s 和 %s 的时间差为 %d 秒\n", timeStr1, timeStr2, timeSub)
	}
}
