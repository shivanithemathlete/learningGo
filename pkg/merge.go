package pkg

func Merge(arr []int, begin, mid, end int) {
	left := mid - begin + 1
	right := end - mid
	leftSub := make([]int, left)
	rightSub := make([]int, right)
	for i := 0; i < left; i++ {
		leftSub[i] = arr[begin+i]
	}
	for i := 0; i < right; i++ {
		rightSub[i] = arr[mid+i+1]
	}
	i, j, index := 0, 0, begin
	for i < left && j < right {
		if leftSub[i] <= rightSub[j] {
			arr[index] = leftSub[i]
			i++
		} else {
			arr[index] = rightSub[j]
			j++
		}
		index++
	}
	for i < left {
		arr[index] = leftSub[i]
		i++
		index++
	}
	for j < right {
		arr[index] = rightSub[j]
		j++
		index++
	}

}
