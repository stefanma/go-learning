package main

import (
	"fmt"
	"sync"
)

func bufferedProducer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		ch <- i // 若缓冲未满，不阻塞
		fmt.Printf("发送 %3d\n", i)

	}
	close(ch) // 发完关闭通道
	fmt.Println("生产者完成，发送了 100 个数")
}

func bufferedConsumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Println("消费:", num)
		fmt.Printf("接收到: %3d\n", num)
	}
}

// func main() {
// 	bufferSize := 10
// 	ch := make(chan int, bufferSize) // 带缓冲通道

// 	var wg sync.WaitGroup
// 	wg.Add(2)

// 	go bufferedProducer(ch, &wg)
// 	go bufferedConsumer(ch, &wg)

// 	wg.Wait()
// }
