package main

func max(numbers []int) int {
	var max int
	for i, number := range numbers {
		if i == 0 || number > max {
			max = number
		}
	}
	return max
}
