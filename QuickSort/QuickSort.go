package main

import (
	"fmt"
)

func doQuickSort(array []int, left int, right int) {
	if left > right {
		return
	}
	l, r := left, right
	base := array[l]
	for l != r {
		for ; array[r] >= base && l < r; r-- {
		}
		for ; array[l] <= base && l < r; l++ {
		}
		if l < r {
			array[l], array[r] = array[r], array[l]
		}
	}
	array[left], array[l] = array[l], array[left]
	doQuickSort(array, left, l-1)
	doQuickSort(array, r+1, right)
	return
}

func quickSort(array []int) {
	doQuickSort(array, 0, len(array)-1)
}

func main() {
	a := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}
	fmt.Println(a)
	quickSort(a)
	fmt.Println(a)
}
