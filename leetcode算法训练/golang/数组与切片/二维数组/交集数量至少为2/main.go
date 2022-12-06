package main

import (
	"fmt"
	"sort"
)

func intersectionSizeTwo(intervals [][]int) int {

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][len(intervals[i])-1] < intervals[j][len(intervals[j])-1]
	})

	s := make([]int, 0, len(intervals))
	var s2 []int

	for line, arr := range intervals {
		//
		if line == 0 {
			s = append(s, arr[len(arr)-2], arr[len(arr)-1])
			//
		} else if arr[len(arr)-1] >= s[len(s)-1] && arr[0] <= s[0] {
			continue
			//
		} else if arr[0] > s[len(s)-1] {
			s = append(s, arr[len(arr)-2], arr[len(arr)-1])
			//
		} else if s[len(s)-1]-arr[0] == 1 {
			s = append(s, arr[len(arr)-1])
			//
		} else if s[len(s)-1]-arr[0] == 0 {
			s = append(s, arr[len(arr)-1])
		} else if s[len(s)-1] == arr[len(arr)-1] && arr[0] > s[0] {

			for i := 0; i < len(arr)-1-1; i++ {
				for j := 0; j < len(s)-1-1; j++ {
					if arr[i] == s[j] {
						continue
					}
				}
			}
			s2 = append(s2, arr[0])
		}
	}
	return len(s) + len(s2)
}

func main() {
	arr := [][]int{
		{1, 2, 3},
		{1, 2, 3, 4},
		{2, 3, 4, 5},
		{3, 4, 5},
	}
	fmt.Println(intersectionSizeTwo(arr))
}
