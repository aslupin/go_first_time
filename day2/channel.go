package main

import "fmt"

func sum(arr []int, c chan int) {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	c <- sum
	// close(c)
}
func main() {
	c := make(chan int)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	go sum(arr[:len(arr)/2], c)
	go sum(arr[len(arr)/2:], c)
	fmt.Println(<-c)
}
