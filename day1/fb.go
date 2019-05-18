package main

import "fmt"

func main() {
	// b := *getPointer()
	// var b *int
	b := *getPointer()
	fmt.Println("Value is", b)
}

func getPointer() (myPointer *int) {
	a := 234
	return &a
}

// func main() {

// 	for i := 1; i <= 100; i++ {

// 		if i%3 == 0 && i%5 == 0 {
// 			fmt.Println("FB")
// 		} else if i%3 != 0 && i%5 == 0 {
// 			fmt.Println("B")
// 		} else if i%3 == 0 && i%5 != 0 {
// 			fmt.Println("F")
// 		} else {
// 			fmt.Println(i)
// 		}
// 	}

// }
