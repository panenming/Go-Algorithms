package main

import "math/rand"

//随机选择一个点作为media
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	median := arr[rand.Intn(len(arr))]
	low := make([]int, 0, len(arr))
	high := make([]int, 0, len(arr))
	middle := make([]int, 0, len(arr))

	for _, item := range arr {
		switch {
		case item < median:
			low = append(low, item)
		case item == median:
			middle = append(middle, item)
		case item > median:
			high = append(high, item)
		}
	}

	low = quickSort(low)
	high = quickSort(high)
	low = append(low, middle...)
	low = append(low, high...)
	return low
}
