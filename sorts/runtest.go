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
	fmt.Println("测试快速排序-----")
	arrayzor = []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
	fmt.Println("unsorted array:", arrayzor)
	arrayzor = quickSort(arrayzor)
	fmt.Println("sorted array:", arrayzor)
	fmt.Println("测试堆栈排序-----")
	arrayzor = []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
	fmt.Println("unsorted array:", arrayzor)
	arrayzor = heapSort(arrayzor)
	fmt.Println("sorted array:", arrayzor)
	fmt.Println("测试希尔排序-----")
	arrayzor = []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
	fmt.Println("unsorted array:", arrayzor)
	shellSort(arrayzor)
	fmt.Println("sorted array:", arrayzor)

	var edges []Edge
	edges = append(edges, Edge{1, 2})
	edges = append(edges, Edge{2, 3})
	edges = append(edges, Edge{4, 3})
	edges = append(edges, Edge{4, 5})
	edges = append(edges, Edge{3, 6})
	edges = append(edges, Edge{5, 6})

	order, levels := topologicalSort(edges)
	fmt.Println("result: ", order)
	fmt.Println("levels: ", levels)

	// or you can specify dependencies
	var deps []Dependency
	deps = append(deps, Dependency{nodeId: 1, children: []int{2}})
	deps = append(deps, Dependency{nodeId: 2, children: []int{3}})
	deps = append(deps, Dependency{nodeId: 4, children: []int{3, 5}})
	deps = append(deps, Dependency{nodeId: 5, children: []int{6}})
	deps = append(deps, Dependency{nodeId: 3, children: []int{6}})

	edges = dependenciesToEdges(deps)
	order, levels = topologicalSort(edges)
	fmt.Println("result: ", order)
	fmt.Println("levels: ", levels)

	fmt.Println("test-----")
	arrayzor = []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
	arr := []int{10, 11}
	arrayzor = append(arrayzor, arr...)
	fmt.Println(arrayzor)
}
