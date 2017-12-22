package main

//使用最大heap的概念，取值的时候需要反向
type MaxHeap struct {
	slice    []int
	heapSize int
}

func buildMaxHeap(slice []int) MaxHeap {
	h := MaxHeap{slice: slice, heapSize: len(slice)}
	for i := len(slice) / 2; i >= 0; i-- {
		h.MaxHeapify(i)
	}
	return h
}

func (h MaxHeap) MaxHeapify(i int) {
	l, r := 2*i+1, 2*i+2
	max := i
	if l < h.size() && h.slice[l] > h.slice[max] {
		max = l
	}
	if r < h.size() && h.slice[r] > h.slice[max] {
		max = r
	}
	if max != i {
		h.slice[i], h.slice[max] = h.slice[max], h.slice[i]
		h.MaxHeapify(max)
	}
}

func (h MaxHeap) size() int {
	return h.heapSize
}

func heapSort(slice []int) []int {
	h := buildMaxHeap(slice)
	for i := h.size() - 1; i >= 1; i-- {
		h.slice[0], h.slice[i] = h.slice[i], h.slice[0]
		h.heapSize--
		h.MaxHeapify(0)
	}
	return h.slice
}
