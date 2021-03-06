package main

import (
	"fmt"
	"time"
)

/*
课后练习 1.2
基于 Channel 编写一个简单的单线程生产者消费者模型：

队列：
队列长度 10，队列元素类型为 int
生产者：
每 1 秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞
消费者：
每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞
*/

func produce(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Printf("%v produced %d\n", time.Now().Format("2006-01-02 15:04:05"), i)
		time.Sleep(1 * time.Second)
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for v := range ch {
		fmt.Printf("%v consumed %d\n", time.Now().Format("2006-01-02 15:04:05"), v)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	ch := make(chan int, 10)

	// produce
	go produce(ch)

	// consumer
	consumer(ch)
}
