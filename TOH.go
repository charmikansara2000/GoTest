package main

import "fmt"

func TOH(n, A, C, B int) {
	/*if n == 1 {
		fmt.Println("Move from ", A, "to ", C)
	}*/
	if n > 0 {
		TOH(n-1, A, B, C)
		fmt.Println("Move disk ", n, "from ", A, "to ", C)
		TOH(n-1, B, C, A)
	}
}
func main() {
	n := 4
	TOH(n, 1, 3, 2)
}
