package Code02_BubbleSort

func bubbleSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	for e := len(arr) - 1; e > 0; e-- {
		for i := 0; i < e; i++ {
			if arr[i] > arr[i+1] {
				swap(arr, i, i+1)
			}
		}
	}
}

// 交换arr的i和j位置上的值
func swap(arr []int, i, j int) {
	arr[i] ^= arr[j]
	arr[j] ^= arr[i]
	arr[i] ^= arr[j]
}
