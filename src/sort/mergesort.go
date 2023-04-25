package sort

func Merge(items []int) []int {
	if len(items) < 2 {
		return items
	}
	mid := len(items) / 2
	a := Merge(items[:mid])
	b := Merge(items[mid:])
	return merge(a, b)
}

func merge(left []int, right []int) []int {
	result := make([]int, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result[i+j] = left[i]
			i++
		} else {
			result[i+j] = right[j]
			j++
		}
	}
	for i < len(left) {
		result[i+j] = left[i]
		i++
	}
	for j < len(right) {
		result[i+j] = right[j]
		j++
	}
	return result
}

// MergeIter  迭代的归并排序
func MergeIter(items []int) []int {
	for step := 1; step < len(items); step += step {
		for i := 0; i+step < len(items); i += 2 * step {
			tmp := merge(items[i:i+step], items[i+step:MinInt(i+2*step, len(items))])
			copy(items[i:], tmp)
		}
	}
	return items
}

func MinInt(values ...int) int {
	min := values[0]
	for _, value := range values {
		if value < min {
			min = value
		}
	}
	return min
}
