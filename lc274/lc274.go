package main

import (
	"fmt"
	"sort"
)

// 思路：排序数组，遍历
// if 第i个元素的值 小于等于 第i个元素后的元素个数（含第i个元素）
//    if 该值等于 第i个元素后的元素个数（含第i个元素）或 第i个元素后的元素个数（不含第i个元素）
//       该值即为h指数
// else 第i个元素的值 大于 第i个元素后的元素个数（含第i个元素） 例如：[100] 或 [1,100,101] 等情况
//		取h指数为 第i个元素后的元素个数（含第i个元素）
func hIndex(citations []int) int {
	sort.Ints(citations)
	for i := 0; i < len(citations); i++ {
		if citations[i] <= len(citations)-i {
			if citations[i] == len(citations)-i-1 || citations[i] == len(citations)-i {
				return citations[i]
			}
		} else {
			return len(citations) - i
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
