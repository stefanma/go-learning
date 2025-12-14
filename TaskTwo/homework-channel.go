package main

import "fmt"

func producer(ch chan<- int) {
	defer close(ch) // 发送完毕后关闭通道
	for i := 1; i <= 10; i++ {
		ch <- i // 发送数据（阻塞直到有接收者）
		fmt.Println("发送:", i)
	}
}

func consumer(ch <-chan int) {
	for num := range ch { // 自动在通道关闭后退出循环
		fmt.Println("接收:", num)
	}
}

// func main() {
// 	ch := make(chan int) // 无缓冲通道

// 	go producer(ch)
// 	// 等待 consumer 处理完所有数据
// 	// 由于 consumer 使用 range，它会在通道关闭后自然退出
// 	// 主 goroutine 需要等待足够时间或使用 sync.WaitGroup
// 	// 这里简单使用 time.Sleep（实际项目建议用 WaitGroup）

// 	// 更健壮的方式：让 main 等待 consumer 完成
// 	// 我们可以用一个 done channel 或 WaitGroup
// 	done := make(chan bool)
// 	// 启动一个 goroutine 等待 consumer 完成
// 	go func() {
// 		consumer(ch)
// 		done <- true
// 	}()

// 	<-done // 等待 consumer 完成
// 	fmt.Println("完成")
// }
