package main

import (
	"fmt"
	"math"
	"sort"
)

func maximumGap(nums []int) int {
	sort.Ints(nums)
	maxGap := 0
	for i := 1; i < len(nums); i++ {
		gap := int(math.Abs(float64(nums[i] - nums[i-1])))
		if gap > maxGap {
			maxGap = gap
		}
	}
	return maxGap
}

func main() {
	array := []int{10, 3, 4, 2, 5, 6, 3, 2, 1, 0}
	fmt.Println(array)
	maximumGap(array)
	fmt.Println(array)
	fmt.Println(maximumGap(array))
}
