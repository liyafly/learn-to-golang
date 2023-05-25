package sort

import (
	"math/rand"
)

func Partition(arr []int, low, high int) int {
	index := low - 1
	pivotElement := arr[high]
	for i := low; i < high; i++ {
		if arr[i] < pivotElement {
			index += 1
			arr[index], arr[i] = arr[i], arr[index]

		}
	}
	arr[index+1], arr[high] = arr[high], arr[index+1]
	return index + 1
}

func PartitionWithRand(arr []int, low, high int) int {
	randIndex := rand.Intn(high-low+1) + low
	arr[randIndex], arr[high] = arr[high], arr[randIndex]
	pivotElement := arr[high]
	index := low - 1
	for i := low; i < high; i++ {
		if arr[i] < pivotElement {
			index += 1
			arr[index], arr[i] = arr[i], arr[index]

		}
	}
	arr[index+1], arr[high] = arr[high], arr[index+1]
	return index + 1
}

func QuicksortRange(arr []int, low, high int) {
	if len(arr) <= 1 {
		return
	}
	if low < high {
		pivot := Partition(arr, low, high)
		QuicksortRange(arr, low, pivot-1)
		QuicksortRange(arr, pivot+1, high)
	}
}

func Quicksort(arr []int) []int {
	QuicksortRange(arr, 0, len(arr)-1)
	return arr
}

// ThreeWayQuicksort 三路快排，优化算法
func ThreeWayQuicksort(nums []int, left, right int) {
	if left >= right {
		return
	}
	// 随机选取基准值
	randIndex := rand.Int()%(right-left+1) + left
	nums[left], nums[randIndex] = nums[randIndex], nums[left]
	pivot := nums[left]
	lt, i, gt := left, left+1, right
	for i <= gt {
		if nums[i] < pivot {
			nums[i], nums[lt] = nums[lt], nums[i]
			lt++
			i++
		} else if nums[i] > pivot {
			nums[i], nums[gt] = nums[gt], nums[i]
			gt--
		} else {
			i++
		}
	}
	ThreeWayQuicksort(nums, left, lt-1)
	ThreeWayQuicksort(nums, gt+1, right)
}

func TwoWayQuickSort(arr []int, left, right int) {
	if left < right {
		pivotIndex := TwoWayPartition(arr, left, right)
		TwoWayQuickSort(arr, left, pivotIndex-1)
		TwoWayQuickSort(arr, pivotIndex+1, right)
	}

}

func TwoWayPartition(arr []int, left, right int) int {
	randIndex := rand.Int()%(right-left+1) + left
	arr[randIndex], arr[left] = arr[left], arr[randIndex]
	pivot := arr[left]
	for left < right {
		for left < right && arr[right] >= pivot {
			right--
		}
		arr[left] = arr[right]
		for left < right && arr[left] <= pivot {
			left++
		}
		arr[right] = arr[left]

	}
	arr[left] = pivot
	return left
}
