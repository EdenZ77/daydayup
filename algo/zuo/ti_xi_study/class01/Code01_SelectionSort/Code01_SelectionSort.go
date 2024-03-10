package Code01_SelectionSort

func selectionSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	// 0 ~ n-1  找到最小值，在哪，放到0位置上
	// 1 ~ n-1  找到最小值，在哪，放到1 位置上
	// 2 ~ n-1  找到最小值，在哪，放到2 位置上
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ { // i ~ N-1 上找最小值的下标
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		swap(arr, i, minIndex)
	}
}

func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
