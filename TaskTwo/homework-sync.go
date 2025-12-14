package main

// func main() {
// 	var counter int
// 	var mu sync.Mutex
// 	var wg sync.WaitGroup

// 	// 启动 10 个协程
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		// 启动一个协程
// 		go func() {
// 			defer wg.Done()
// 			for j := 0; j < 1000; j++ {
// 				mu.Lock()
// 				counter++
// 				mu.Unlock()
// 			}
// 		}()
// 	}
// 	// 等待所有协程结束
// 	wg.Wait()
// 	fmt.Println("最终计数器值（Mutex）:", counter)
// }
