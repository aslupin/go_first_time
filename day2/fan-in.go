package main

import (
	"fmt"
	"time"
)

func worker2(c chan int) {
	// for {
	time.Sleep(1 * time.Second)
	c <- 10
	// }

}
func worker1(c chan int) {
	// for {
	time.Sleep(4 * time.Second)
	c <- 20
	// }

}

func worker3(c chan int) {
	// for {
	time.Sleep(2 * time.Second)
	fmt.Println(<-c)
	// }
}

func timer(stop chan int) {
	time.Sleep(5 * time.Second)
	stop <- 1
}
func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)
	stop := make(chan int)
	go worker1(c1)
	go worker2(c2)
	go worker3(c3)
	go timer(stop)
	// n1 := <-c1
	// n2 := <-c2
LOOP:
	for {
		select {
		case n := <-c1:
			fmt.Println(n)
		case n := <-c2:
			fmt.Println(n)
		case c3 <- 10:
			fmt.Println("send to c3")
		case <-stop:
			break LOOP
		}
	}

}
