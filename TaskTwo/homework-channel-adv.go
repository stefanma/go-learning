package main

import (
	"fmt"
	"sync"
)

// 生产者：发送 1~10
func producerAdv(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		ch <- i // 无缓冲通道：阻塞直到消费者接收
		fmt.Println("发送:", i)
	}
	close(ch) // 发送完毕，关闭通道
}

// 消费者：接收并打印
func consumerAdv(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch { // 自动在通道关闭后退出
		fmt.Println("接收到:", num)
	}
}

// func main() {
// 	ch := make(chan int) // 无缓冲通道
// 	var wg sync.WaitGroup

// 	wg.Add(2)
// 	go producerAdv(ch, &wg)
// 	go consumerAdv(ch, &wg)

// 	wg.Wait() // 等待两个协程完成
// }
