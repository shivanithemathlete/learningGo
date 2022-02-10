package pkg

func Reduce(x []int, f func(x int) int) []int {
	var res []int
	for i := range x {
		res = append(res, f(x[i]))
	}
	return res
}
