package sort

func HeapSort(arr []int) {
	n := len(arr)
	// 构造大根堆
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}
	// 从堆顶开始，依次将最大值防盗序列末尾
	for i := n - 1; i >= 0; i-- {
		// 将堆顶元素与末尾元素交换
		arr[0], arr[i] = arr[i], arr[0]
		// 重新构造大根堆
		heapify(arr, i, 0)
	}
}

// heapify
func heapify(arr []int, n, i int) {
	largest := i     // 先假设最大值为当前位置
	left := 2*i + 1  // 左节点位置
	right := 2*i + 2 // 右节点位置
	// 如果左子节点比当前节点大，则更新最大值位置
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	// 如果右子节点比最大值大，则更新最大值位置
	if right < n && arr[right] > arr[largest] {
		largest = right
	}
	// 如果最大值不是当前位置，则需要交换位置，并递归调整子树
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}
