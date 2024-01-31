package arrays

func Sum(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}

func SumAll(numToSum ...[]int) []int {
	var sums []int

	// sums := make([]int, lenOfNumbers)

	for _, numbers := range numToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numToSum ...[]int) []int {
	var sums []int

	// sums := make([]int, lenOfNumbers)

	for _, numbers := range numToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}

	}

	return sums
}
