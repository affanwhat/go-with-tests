package arrays

func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num 
	}
	return sum
}

func SumAll(numberstoSum ...[]int) []int {
	lengthofNumbers := len(numberstoSum)
	sums := make([]int, lengthofNumbers)
	for i, numbers := range numberstoSum {
		sums[i] = Sum(numbers)
	}
	return sums
}

/*
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
*/

func SumAllTails(numberstoSum ...[]int) []int {
	lengthofNumbers := len(numberstoSum)
	sums := make([]int, lengthofNumbers)
	for i, numbers := range numberstoSum {
		if len(numbers) < 1 {
			sums[i] = 0
		} else {
			sums[i] = Sum(numbers[1:])
		}
	}
	return sums
}