package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func SumAll(numbers [][]int, excludeHead bool) []int {
	var sum []int
	for _, arr := range numbers {
		currSum := Sum(arr)
		if excludeHead {
			currSum -= arr[0]
		}
		sum = append(sum, currSum)
	}
	return sum
}
