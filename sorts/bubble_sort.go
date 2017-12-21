package main

// 比较相邻的两个数字，如果排在前边的数字大于排在后边的就交换两个的位置，直到没有可以交换的
func bubbleSort(arrayzor []int) {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(arrayzor)-1; i++ {
			if arrayzor[i+1] < arrayzor[i] {
				swap(arrayzor, i, i+1)
				swapped = true
			}
		}
	}
}

func swap(arrayzor []int, i, j int) {
	arrayzor[j], arrayzor[i] = arrayzor[i], arrayzor[j]
}
