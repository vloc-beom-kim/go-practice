package main

import (
	"fmt"
	"math"
)

func min(nums ...int) {
	min := math.MaxInt64
	for _, v := range nums {
		if min > v {
			min = v
		}
	}
	fmt.Println(min)
}

func main() {
	nums := []int{5, 1, 7, 8, 3, 9}
	min(nums...)
}
