package sorting

func selectionSort(arr []int) {
	for i := len(arr) - 1; i > 0; i-- {
		maxVal := i
		for j := i - 1; j >= 0; j-- {
			if arr[j] < arr[maxVal] {
				maxVal = j
			}
		}
		arr[maxVal], arr[i] = arr[i], arr[maxVal]
	}
}
