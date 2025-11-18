package goroutine

import (
	"fmt"
	"sync"
	"time"
)

var (
	count  int
	mutex  sync.Mutex
	wgSync sync.WaitGroup
)

func increment() {
	defer wgSync.Done() // 任务完成, 计数器减1
	for i := 0; i < 1000; i++ {
		mutex.Lock()   // 加锁,独占资源
		count++        // 临界区操作(修改共享变量)
		mutex.Unlock() // 解锁,释放资源
		time.Sleep(1 * time.Microsecond)
	}
}

func SyncMutex() {
	fmt.Println("=====syncmutex======")
	wgSync.Add(2) // 启动两个 Goroutine, 计数器加2

	go increment()
	go increment()

	wgSync.Wait()
	fmt.Printf("最终 count 值: %d\n", count)

}