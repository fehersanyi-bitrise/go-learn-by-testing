package arrays

func sum(numbers []int) int {
	s := 0
	for _, number := range numbers {
		s += number
	}
	return s
}

func sumAll(numbersToSum ...[]int) (sums []int) {
	for _, slices := range numbersToSum {
		sums = append(sums, sum(slices))
	}
	return
}

func sumAllTails(slicesToSum ...[]int) (tails []int) {
	for _, slices := range slicesToSum {
		if len(slices) == 0 {
			tails = append(tails, 0)
		} else {
			tails = append(tails, sum(slices[1:]))
		}
	}
	return
}
