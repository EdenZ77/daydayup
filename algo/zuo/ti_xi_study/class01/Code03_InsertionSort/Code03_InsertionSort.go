package Code03_InsertionSort

func insertionSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
			swap(arr, j, j+1)
		}
	}
}

func swap(arr []int, i, j int) {
	if i == j {
		return // 如果i和j是同一位置，则不执行swap操作
	}
	arr[i] = arr[i] ^ arr[j]
	arr[j] = arr[i] ^ arr[j]
	arr[i] = arr[i] ^ arr[j]
}
