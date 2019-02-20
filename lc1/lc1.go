package main

import (
	"fmt"
)

// 两次遍历
func twoSumOn2(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i]+nums[j] == target && i != j {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}

// 一次遍历数组存入Map，二次遍历数组查找Map
func twoSumHash(nums []int, target int) []int {
	numsMap := make(map[int]int, len(nums))
	for k, v := range nums {
		numsMap[v] = k
	}
	for i := range nums {
		index, ok := numsMap[target-nums[i]]
		if ok && i != index {
			return []int{i, index}
		}
	}
	return []int{-1, -1}
}

// 一次遍历数组存入Map，同时查找Map已有元素
func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int, len(nums))
	for k, v := range nums {
		index, ok := numsMap[target-nums[k]]
		if ok && k != index {
			return []int{k, index}
		}
		numsMap[v] = k
	}
	return []int{-1, -1}
}

func main() {
	nums := []int{3, 9, 6, 1, 2, 5}
	// nums := []int{3, 3}
	target := 6
	ans := twoSum(nums, target)
	fmt.Println(ans)
}
