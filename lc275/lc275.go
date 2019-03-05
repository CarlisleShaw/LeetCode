package main

import "fmt"

func hIndex(citations []int) int {
	left := 0
	right := len(citations) - 1
	for left < right {
		mid := left + ((right - left) >> 1)
		if citations[mid] >= len(citations)-mid {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if citations[right] < len(citations)-right {
		return 0
	}
	return len(citations) - right
}

func main() {
	// citations := []int{0, 1, 3, 5, 6}
	citations := []int{0, 1, 2, 4, 100, 101, 102, 103, 104, 105, 106, 107}
	// citations := []int{0, 1, 2, 5, 6, 100}
	ans := hIndex(citations)
	fmt.Println(ans)
}
