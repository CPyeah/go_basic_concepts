package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup;

func main()  { 
	
	for i := 0; i < 10; i++ {
		url := "url_" + string(rune(i+'0'));
		wg.Add(1)
		go download(url);	 
	}
	wg.Wait();
	fmt.Println("done")
	
}

func download(url string) {
	fmt.Println(url)
	time.Sleep(time.Second)
	wg.Done();
}