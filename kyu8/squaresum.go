package kata

func SquareSum(numbers []int) int {
	squ := 0
	for _, num := range numbers {
		squ += num * num
	}
	return squ
}
