package arraysandslices

import "slices"

func Sum(numbers []int) int {
	var sum int
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	length := len(numbersToSum)
	sum := make([]int, length)
	for i, v := range numbersToSum {
		sum[i] = Sum(v)
	}
	return sum
}

func SumAllTails(numbersToSum ...[]int) []int {
	length := len(numbersToSum)
	sum := make([]int, length)
	var tail []int
	for i, numbers := range numbersToSum {
		if slices.Equal(numbers, []int{}) {
			tail = []int{}
		} else {
			tail = numbers[1:]
		}
		sum[i] = Sum(tail)
	}
	return sum
}
