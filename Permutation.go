package main

import (
	"fmt"
	"reflect"
)

func permutation(s []int, n int) {
	//var res []int
	count := 0
	swapF := reflect.Swapper(s)
	for i := 0; i < len(s); i++ {
		lock := s[i]

		swapF(0, i)
		var temp []int
		temp = append(temp, lock)

		for j := 1; j < len(s); j++ {
			temp = append(temp, s[j])
		}

		if len(temp) == 3 {
			count++
			if count == n {
				fmt.Println(temp)
				defer fmt.Println("Answer:", temp)
			} else {
				fmt.Println(temp)
			}

			swapT := reflect.Swapper(temp)
			swapT(1, 2)

			count++
			if count == n {
				fmt.Println(temp)
				defer fmt.Println("Answer:", temp)
			} else {
				fmt.Println(temp)
			}
		}

	}

}
func main() {
	str := []int{1, 2, 3}
	n := 2
	permutation(str, n)

}
