package main

//每次将剩余的插入已经排序好的数组
func insertionSort(arr []int) {
	for out := 1; out < len(arr); out++ {
		temp := arr[out]
		in := out
		for ; in > 0 && arr[in-1] >= temp; in-- {
			arr[in] = arr[in-1]
		}
		arr[in] = temp
	}
}
