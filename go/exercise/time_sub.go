package main

import (
	"fmt"
	"time"
)

func timeSub(timeStr1, timeStr2 string) (int, error) {
	// 定义时间布局（必须严格对应时间字符串的格式）
	// 布局规则：用固定参考时间 2006/01/02 15:04:05 对应你的时间格式
	layout := "2006/01/02 15:04:05" // 时间格式模板

	// 1. 解析时间字符串为 time.Time 类型
	time1, err := time.Parse(layout, timeStr1)
	if err != nil {
		fmt.Printf("解析时间 %s 失败: %v\n", timeStr1, err)
		return -1, nil
	}
	time2, err := time.Parse(layout, timeStr2)
	if err != nil {
		fmt.Printf("解析时间 %s 失败: %v\n", timeStr2, err)
		return -1, nil
	}

	// 2. 计算时间差（确保用晚的时间减早的时间，避免负数值）
	var timeSub time.Duration
	if time1.After(time2) {
		timeSub = time1.Sub(time2)
	} else {
		timeSub = time2.Sub(time1)
	}

	// 3. 转换为秒数（两种方式：浮点数/整数）
	// result := timeSub.Seconds() // 保留小数
	result := int(timeSub.Seconds())

	return result, nil
}

