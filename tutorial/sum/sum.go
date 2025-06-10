package sum

func Sum(numbers []int) int {
	var result int = 0
	for _, num := range numbers {
		result += num
	}
	return result
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sum:= Sum(numbers);
		sums = append(sums, sum)
	}
	return sums
}