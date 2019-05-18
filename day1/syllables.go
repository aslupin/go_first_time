package main

import "fmt"

func main() {
	aeiou := [5]string{"a", "e", "i", "o", "u"}
	var n int
	var txt string
	fmt.Scanf("%v\n", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%v\n", &txt)
		isSide := false
		for j := 0; j < len(txt); j++ {
			in, val := in_array(txt[i], aeiou)
			if in {

			}
		}
	}

}
