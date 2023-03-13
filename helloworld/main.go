package main

import "fmt"

func main() {
	s := [11]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, ele := range s {
		if ele%2 != 0 {
			fmt.Println(i, " is Odd")
		} else {
			fmt.Println(i, " is Even")
		}
	}
}
