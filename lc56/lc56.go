package main

import (
	"fmt"
	"sort"
)

// Interval 区间
type Interval struct {
	Start int
	End   int
}

// Axis 数轴节点
type Axis struct {
	Value int
	Start bool // true为起点，false为终点
}

type axisArray []Axis

func (a axisArray) Len() int           { return len(a) }
func (a axisArray) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a axisArray) Less(i, j int) bool { return a[i].Value < a[j].Value }

// 将各区间的起点和终点统一放置在数轴上
// 实现方式：提取Interval数组中的起点和终点，存入数轴数组，
// 对数轴数组进行*稳定排序*（否则，例如[[0,0],[1,2]]情况下，排序后会出现0值终点排在0值起点前的情况）
func toAxis(intervals []Interval) axisArray {
	ans := make(axisArray, len(intervals)*2)
	i := 0
	for _, v := range intervals {
		ans[i] = Axis{v.Start, true}
		ans[i+1] = Axis{v.End, false}
		i += 2
	}
	// 稳定排序
	sort.Stable(ans)
	return ans
}

// 思路：
// 将各区间的起点和终点统一放置在数轴上，
// 定义计数器 count := 0，遍历数轴数组，
// 遇到起点 count++，遇到终点 count--，实现合并区间
func merge(intervals []Interval) []Interval {
	axis := toAxis(intervals)
	count := 0
	start := 0
	for i := 0; i < len(axis); i++ {
		if axis[i].Start {
			if count == 0 {
				// 记录起始节点
				start = i
			}
			count++
		} else {
			if count-1 == 0 {
				// 考虑例如：[[1,2],[2,3]] 的情况
				if i+1 < len(axis)-1 && axis[i].Value == axis[i+1].Value {
					i++
					continue
				}
				// 去除中间节点
				axis = append(axis[:start+1], axis[i:]...)
				i = start + 1 // 本次循环执行结束时还会执行 i++
			}
			count--
		}
	}
	// 从数轴数组 转回 Interval数组
	ans := make([]Interval, len(axis)/2)
	for i := 0; i < len(axis)/2; i++ {
		ans[i].Start = axis[2*i].Value
		ans[i].End = axis[2*i+1].Value
	}
	return ans
}

func main() {
	// i := []Interval{{0, 0}, {3, 3}, {1, 1}, {1, 3}, {2, 6}, {8, 10}, {15, 18}, {-1, -1}}
	// i := []Interval{{1, 4}, {4, 5}, {4, 4}, {6, 6}}
	// i := []Interval{{2, 3}, {2, 2}, {3, 3}, {1, 3}, {5, 7}, {2, 2}, {4, 6}}
	i := []Interval{{0, 0}, {0, 2}, {3, 4}, {1, 1}, {1, 2}, {2, 3}, {4, 5}}
	// i := []Interval{{2, 3}, {5, 5}, {2, 2}, {3, 4}, {3, 4}}
	fmt.Println(i)
	ans := merge(i)
	fmt.Println(ans)
}
