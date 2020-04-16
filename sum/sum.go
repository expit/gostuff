package main

// Sum sums the sum
func Sum(numbers []int) int {
	sum := 0
	for _, i := range numbers {
		sum += i
	}
	return sum
}

// func main() {
// 	fmt.Println("do not use dat function!\nNevur Evur Evur!")
// }

// SumAll sums all the slices in the slice of slices
func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return
}

// SumAllTails sums all the slices in the slice of slices
func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(numbers[1:]))
		}
	}

	return
}
