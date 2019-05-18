package main

import (
	"fmt"
	"strings"
)

func minus(a int, b int) int {
	return a - b
}

func main() {
	c := make(chan string)
	store := make(map[string]string)
	var command string
	var tempKey, tempValue string
	for {
		fmt.Printf("> ")
		fmt.Scanf("%v %v %v", &command, &tempKey, &tempValue)
		if command == "GET" {
			_, ok := store[tempKey]
			if ok {
				fmt.Println(store[tempKey])
			} else {
				fmt.Printf("Not Found\n")
			}

		} else if command == "SET" {
			store[tempKey] = tempValue
			fmt.Printf("OK\n")
		} else if command == "MGET" {
			keys := strings.Split(tempKey, ",")
			wg.Add(len(keys))
			for i := 0; i < len(keys); i++ {
				go func() {
					_, ok := store[keys[i]]
					if ok {
						fmt.Println(keys[i], store[keys[i]])
					} else {
						fmt.Printf("Not Found\n")
					}
				}()
				wg.Done()
			}
			wg.Wait()
			// for i := 0; i < len(keys); i++ {
			// 	_, ok := store[keys[i]]
			// 	if ok {
			// 		fmt.Println(keys[i], store[keys[i]])
			// 	} else {
			// 		fmt.Printf("Not Found\n")
			// 	}
			// }

		}

	}
}
