package main

// 递归版的归并排序不适合数组很大的情况
func merge(a []int, b []int) []int {
	r := make([]int, len(a)+len(b))
	var i, j = 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			r[i+j] = a[i]
			i++
		} else {
			r[i+j] = b[j]
			j++
		}
	}
	if i < len(a) {
		r = append(r[:i+j], a[i:]...)
	}
	if j < len(b) {
		r = append(r[:i+j], b[j:]...)
	}

	return r
}

func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	middle := len(items) / 2
	a := mergeSort(items[:middle])
	b := mergeSort(items[middle:])
	return merge(a, b)
}
