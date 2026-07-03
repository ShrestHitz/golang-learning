package main

import "fmt"

func sumAll(nums ...int) int {
	total := 0

	for _, currentValue := range nums {
		total += currentValue
	}

	return total
}

func main() {

	fmt.Println(sumAll(1, 2, 3, 4, 5))

	values := []int{10, 20}
	fmt.Println(sumAll(values...))

	res := func(n int) int {
		return n * 2
	}

	fmt.Println(res(2))

	// IIFE - immediately invoked function expression
	result1 := func(a int, b int) int {
		return a + b
	}(5, 10) //iife style

	fmt.Println((result1))
}
