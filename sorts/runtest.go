package main

import (
	"fmt"
)

func main() {
	fmt.Println("测试冒泡排序-----")
	arrayzor := []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
	fmt.Println("unsorted array:", arrayzor)
	bubbleSort(arrayzor)
	fmt.Println("sorted array:", arrayzor)
	fmt.Println("测试插入排序-----")
	arrayzor = []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
	fmt.Println("unsorted array:", arrayzor)
	insertionSort(arrayzor)
	fmt.Println("sorted array:", arrayzor)
	fmt.Println("测试选择排序-----")
	arrayzor = []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
	fmt.Println("unsorted array:", arrayzor)
	selectionSort(arrayzor)
	fmt.Println("sorted array:", arrayzor)
	fmt.Println("测试归并排序-----")
	arrayzor = []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
	fmt.Println("unsorted array:", arrayzor)
	arrayzor = mergeSort(arrayzor)
	fmt.Println("sorted array:", arrayzor)

	fmt.Println("test-----")
	arrayzor = []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
	arr := []int{10, 11}
	arrayzor = append(arrayzor, arr...)
	fmt.Println(arrayzor)
}
