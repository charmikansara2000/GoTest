package main

import (
	"fmt"
	"math"
)

func JumpGame(nums []int) int {
	i := 0
	count := 0
	//var max_index int
	max_index := nums[i]
	for i < len(nums) && i <= max_index {
		new_index := nums[i] + i
		if new_index > max_index {
			count++
		}
		max_index = int(math.Max(float64(new_index), float64(max_index)))
		i += 1

	}
	if i != len(nums) {
		fmt.Println("can not reach at the end")
	}
	return count
}
func main() {
	nums := []int{1, 1, 2, 0, 1, 4}
	jumps := JumpGame(nums)
	fmt.Println("No. of jumps:", jumps)
}

