package main

import (
	"fmt"
)

func worker(buff chan int) {

	fmt.Println(<-buff)
	fmt.Println(<-buff)
	fmt.Println(<-buff)
}

func main() {
	bufferChannel := make(chan int, 3)
	go worker(bufferChannel)
	bufferChannel <- 10
	bufferChannel <- 11
	bufferChannel <- 12

	// fmt.Printf(<- bufferChannel)
	// fmt.Printf(<- bufferChannel)

}
