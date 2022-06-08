package chapter4

func SumArray(numbers [5]int) int {
	sums := 0
	for i := 0; i < 5; i++ {
		sums += numbers[i]
	}
	return sums
}

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			// slice[low: high], this will take value from 1 to end
			tails := numbers[1:]
			sums = append(sums, Sum(tails))
		}
	}
	return sums
}
