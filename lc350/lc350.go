package main

import (
	"fmt"
	"sort"
)

func toMap(nums []int) map[int]int {
	var numsMap map[int]int
	numsMap = make(map[int]int)
	numsMap[nums[0]] = 0
	key := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] == key {
			numsMap[key]++
		} else {
			numsMap[nums[i]] = 1
			key = nums[i]
		}
	}
	return numsMap
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}
	sort.Ints(nums1)
	sort.Ints(nums2)
	nums1map := toMap(nums1)
	nums2map := toMap(nums2)
	ansMap := make(map[int]int)
	count := 0
	for k1, v1 := range nums1map {
		for k2, v2 := range nums2map {
			if k1 == k2 {
				ansMap[k1] = min(v1, v2)
				count += ansMap[k1]
			}
		}
	}
	ans := make([]int, count)
	count = 0
	for k, v := range ansMap {
		for i := 0; i < v; i++ {
			ans[count] = k
			count++
		}
	}
	return ans
}

func main() {
	// nums1 := []int{1, 2, 2, 1}
	// nums2 := []int{2, 2}
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	// nums1 := []int{}
	// nums2 := []int{}
	ans := intersect(nums1, nums2)
	fmt.Println(nums1)
	fmt.Println(nums2)
	fmt.Println(ans)
}
