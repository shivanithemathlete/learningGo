package pkg

func MergeSort(arr []int, begin, end int) {

	if begin >= end {
		return
	}
	mid := begin + (end-begin)/2
	MergeSort(arr, begin, mid)
	MergeSort(arr, mid+1, end)
	Merge(arr, begin, mid, end)

}
