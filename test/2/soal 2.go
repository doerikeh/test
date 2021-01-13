package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 33, 44, 55, 33, 44, 22, 44, 33, 2, 77, 66, 1, 2, 3, 4, 5, 6, 7, 89, 3, 3, 8, 9, 75, 4, 3, 2, 2, 1, 3}
	limit := 11

	for i := 0; i < len(arr); i += limit {
		batch := arr[i:min(i+limit, len(arr))]
		sort.Sort(sort.Reverse(sort.IntSlice(batch)))
		fmt.Println(batch)
	}

	for j := 0; j < len(arr); j += limit {
		batch := arr[j:min(j+limit, len(arr))]
		fmt.Println("result Sum :", sumarray(batch))
	}

	for l := 0; l < len(arr); l += limit {
		batch := arr[l:min(l+limit, len(arr))]
		sum := sumarray(batch)
		avg := (float64(sum)) / (float64(limit))
		fmt.Println("result Avg :", avg)
	}

	for k := 0; k < len(arr); k += limit {
		batch := arr[k:min(k+limit, len(arr))]
		min, max := findMinAndMax(batch)
		fmt.Println("Max", max)
		fmt.Println("Min", min)

	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func sumarray(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func findMinAndMax(a []int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}
