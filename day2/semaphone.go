package main

import (
	"fmt"
	"time"
)

func workerSe(sum chan int, job int) {
	time.Sleep(2 * time.Second)
	fmt.Println(sum, job)
	<-sum

}

func main() {
	sum := make(chan int, 4)
	for i := 0; i < 10000; i++ {
		sum <- i
		go workerSe(sum, i)
	}

}
