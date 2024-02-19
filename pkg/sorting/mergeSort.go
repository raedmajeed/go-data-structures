package sorting

import "fmt"

func SortMain() {
	arr := []int{1, 4, 1, 6, 3, 6, 9}
	MergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr) // uncomment for merge sort
}

func MergeSort(arr []int, left, right int) {
	if left < right {
		mid := (left + right) / 2
		MergeSort(arr, left, mid)
		MergeSort(arr, mid+1, right)
		merge(arr, left, right, mid)
	}
}

func merge(arr []int, left, right, mid int) {
	temp := make([]int, right-left+1)
	i, j := left, mid+1
	k := 0
	for i <= mid && j <= right {
		if arr[i] < arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			j++
		}
		k++
	}
	for i <= mid {
		temp[k] = arr[i]
		i++
		k++
	}

	for j <= right {
		temp[k] = arr[j]
		j++
		k++
	}

	for i, k := left, 0; i <= right; i, k = i+1, k+1 {
		arr[i] = temp[k]
	}
}
