package ex1

func CreateSlice(size int) []int {
	slice := make([]int, size)
	for i := range size {
		slice[i] = i
	}
	return slice
}
