package goroutine

import (
	"fmt"
	"time"
)

// 并行执行的函数
func task(name string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("任务 %s：执行第 %d 次\n", name, i)
		time.Sleep(100 * time.Millisecond) // 模拟任务耗时
	}
}

func BaseTest() {
	fmt.Println("======goroutine======")
	// 启动 2 个 Goroutine 并行执行
	go task("A")
	go task("B")

	// 主线程等待 Goroutine 执行完成
	// 任务耗时不确定，会导致 “等待过久” 或 “提前退出”
	time.Sleep(500 * time.Millisecond) // 简单等待（不推荐，仅用于演示）
	fmt.Println("主线程结束")
}
