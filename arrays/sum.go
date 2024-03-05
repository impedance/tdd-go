package main

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sumNumbers []int

	for _, numbers := range numbersToSum {
    if len(numbers) == 0 {
		  sumNumbers = append(sumNumbers, 0)
    } else {
      tail := numbers[1:]
      sumNumbers = append(sumNumbers, Sum(tail))
    }
	}

	return sumNumbers
}

// func SumAllTails(numbersToSum ...[]int) []int {
//   return nil
// }
