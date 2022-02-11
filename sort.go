package main

import "fmt"

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

func MergeSort(arr []int, begin, end int) {

	if begin >= end {
		return
	}
	mid := begin + (end-begin)/2
	MergeSort(arr, begin, mid)
	MergeSort(arr, mid+1, end)
	Merge(arr, begin, mid, end)

}

func main() {
	fmt.Println("Using Merge Sort : ")
	arr1 := []int{4, 3, 6, 0, -2, 8, 10}
	fmt.Println(arr1)
	fmt.Println("Initial Array : ", arr1)
	MergeSort(arr1, 0, len(arr1)-1)
	fmt.Println("After Sorting : ", arr1)

}
