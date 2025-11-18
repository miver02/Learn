package goroutine

import "fmt"


func calculate(a, b int, resChan chan int) {
	res := a + b
	resChan <- res // 发送结果到Channel
}

func ChanInt() {
	fmt.Println("======Channel传值=======")
	resChan := make(chan int, 2) // 带缓冲channel(容量2, 避免堵塞)

	// 启动2个, Goroutine并行计算 先进后出
	go calculate(30, 40, resChan)
	go calculate(10, 20, resChan)

	// 接受结果(会堵塞, 知道有数据发送)
	res1 := <-resChan
	res2 := <-resChan

	fmt.Printf("结果1: %d, 结果2: %d, 总和: %d\n", res1, res2, res1+res2)
	close(resChan)
}