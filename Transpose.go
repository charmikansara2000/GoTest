package main

import "fmt"

func Transpose(a [][]int) {
	var trans [3][3]int
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			trans[i][j] = a[j][i]
		}
	}
	for i := 0; i < len(trans); i++ {
		for j := 0; j < len(trans[0]); j++ {
			fmt.Print(trans[i][j])
		}
		fmt.Println("")
	}

}
func main() {
	a := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	Transpose(a)

}
