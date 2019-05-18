package main

import "fmt"

func gen() chan int {
	output := make(chan int)
	go func() {
		for i := 1; i < 10; i++ {
			output <- i
		}
		close(output)
	}()
	return output

}
func mul(input chan int) chan int {
	output := make(chan int)
	go func() {
		for num := range input {
			output <- num * num
		}
		close(output)
	}()
	return output
}

func sumd(input chan int) chan int {
	output := make(chan int)
	go func() {
		for num := range input {
			output <- num + num
		}
		close(output)
	}()
	return output
}

func print(input chan int) {
	for num := range input {
		fmt.Println(num)
	}
}

func main() {
	g := gen()
	m := mul(g)
	s := sumd(m)
	print(s)
}
