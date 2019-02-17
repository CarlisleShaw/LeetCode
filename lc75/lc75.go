package main
import "fmt"

// 思路一：
// 遍历1次，把遇到的0往最左边放，2往最右边放

func sortColors(nums []int)  {
    left := 0
    right := len(nums) - 1
    for i := 0; i < len(nums); i++ {
    	if nums[i] == 0 {
    		if i > left {
    			nums[i], nums[left] = nums[left], nums[i]
    			left++
    			i--		// 对调换到该位置的数进行检查
    		}
    	} else if nums[i] == 2 {
    		if i < right {
    			nums[i], nums[right] = nums[right], nums[i]
    			right--
    			i--
    		}
    	} else if nums[i] == 1 {}
    } 
}

// 思路二：
// 遍历2次，第一次记录0和1的个数，第二次根据0和1的个数按次序重写数组

func sortColors2(nums []int)  {
	zero := 0
	one := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zero++
		} else if nums[i] == 1 {
			one++
		} else {}
	}
	i := 0
	for ; i < zero; i++ {
		nums[i] = 0
	}
	for ; i < zero + one; i++ {
		nums[i] = 1
	}
	for ; i < len(nums); i++ {
		nums[i] = 2
	}
}

func main() {
	a := []int{2,0,1,2,0,2,1,1,0,2,1,2,1,0,2,1,0,0,1}
	sortColors(a)
	fmt.Println(a)
}