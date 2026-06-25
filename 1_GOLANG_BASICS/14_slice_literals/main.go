package main

import (
	"fmt"
)

func main() {
	// common collection type
	// dynamic and can grow
	// []type{...}

	result := []string{"Shresth", "John"}
	fmt.Println(result, result[0], result[len(result)-1])

	result[1] = "Priya"
	fmt.Println(result)

	var nums []int

	nums = append(nums, 10)
	nums = append(nums, 20)
	fmt.Println(nums)
}
