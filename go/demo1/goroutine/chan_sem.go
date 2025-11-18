package goroutine

import (
	"fmt"
	"sync"
	"time"
)


var wgSem sync.WaitGroup

func semTask(id int, sem chan struct{}) {
	defer wgSem.Done()
	defer func() { <-sem }() // 任务完成, 释放信号量
	fmt.Printf("任务 %d 开始执行\n", id)
	time.Sleep(500 * time.Millisecond) // 模拟耗时
	fmt.Printf("任务 %d 执行完成\n", id)
}

func ChanSem() {
	fmt.Println("======channel作为信号量=====")

	const maxParallel = 2
	// wg.Add(2)
	sem := make(chan struct{}, maxParallel) // 信号量Channel

	// 启动 5 个任务, 最多 2 个并行执行
	for i := 1; i <= 5; i++ {
		sem <- struct{}{} // 申请信号量(满了会堵塞,限制并行数)
		wgSem.Add(1)
		go semTask(i, sem)
	}

	// 等待所有任务完成（简单等待，实际可用 WaitGroup）
	// time.Sleep(3 * time.Second)
	wgSem.Wait() // 等待所有Goroutine完成
	close(sem)
	fmt.Println("所有任务结束")
}