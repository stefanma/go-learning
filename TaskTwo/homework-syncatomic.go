package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64 // 必须是 int64（或 uint64 等）
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1)
				fmt.Println("atomic:", counter)
			}
		}()
	}

	wg.Wait()
	fmt.Println("最终计数器值（Atomic）:", counter)
}
