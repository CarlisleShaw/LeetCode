package main

import (
	"fmt"
	"sort"
)

type strArray []string

func (a strArray) Len() int      { return len(a) }
func (a strArray) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a strArray) Less(i, j int) bool {
	if len(a[i]) == len(a[j]) {
		for x := 0; x < len(a[i]); x++ {
			if a[i][x] == a[j][x] {
				continue
			}
			return a[i][x] < a[j][x]
		}
	}
	return len(a[i]) > len(a[j])
}

func longestWord(words []string) string {
	sort.Sort(strArray(words))
	// each word
	for i, str := range words {
		count := 0
		// 遍历每一个切片
		for j := 1; j < len(str); j++ {
			// 从后一个开始遍历查找与切片相等的string
			for _, str2 := range words[i+1:] {
				if str2 == str[:j] {
					count++
					break
				}
			}
			if count != j {
				break
			}
		}
		if count+1 == len(str) {
			return str
		}
	}
	return ""
}

func main() {
	// words := []string{"w", "wo", "wor", "worl", "world"}
	// words := []string{"a", "banana", "app", "appl", "ap", "apply", "apple"}
	words := []string{"yo", "ew", "fc", "zrc", "yodn", "fcm", "qm", "qmo", "fcmz", "z", "ewq", "yod", "ewqz", "y"}
	fmt.Println(words)
	ans := longestWord(words)
	fmt.Println(words)
	fmt.Println(ans)
}
