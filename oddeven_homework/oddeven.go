package main

import "fmt"

type nums []int

func (n nums) findOddEven() {
	for i := range n {
		if n[i]%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}
}
