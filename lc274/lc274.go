package main

import (
	"fmt"
	"sort"
)

func hIndex(citations []int) int {
	sort.Ints(citations)
	for i := 0; i < len(citations); i++ {
		if citations[i] <= len(citations)-i {
			if citations[i] == len(citations)-i-1 || citations[i] == len(citations)-i {
				return citations[i]
			}
		} else {
			for j := 1; (i > 0 && citations[i]-j >= citations[i-1]) || i == 0; j++ {
				if citations[i]-j == len(citations)-i {
					return citations[i] - j
				}
			}
		}
	}
	return 0
}

func main() {
	citations := []int{3, 0, 6, 1, 5, 7, 8, 9, 4, 10}
	// citations := []int{2, 3, 4, 5, 6}
	// citations := []int{100}
	ans := hIndex(citations)
	fmt.Println(citations)
	fmt.Println(ans)
}
