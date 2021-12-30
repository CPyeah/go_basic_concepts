package main

import (
	"fmt"
	"time"
)

var ch = make(chan string, 100);
func main()  { 
	producer();
	consumer();
}

// 生产者
func producer()  {
	for i := 0; i < 10; i++ {
		go produce(string(i+'0'))
	}
}

func produce(s string) {
	time.Sleep(time.Second)
	fmt.Println("produce " + s)
	ch <- s;
}

// 消费者
func consumer() {
	for i := 0; i < 10; i++ {
		r := <-ch;
		fmt.Println("custom " + r)
	}
}