package main

import (
	"fmt"
)

func main() {
	//make([]T, len, cap)- useful for performance as it avoids repetitions
	scores := make([]int, 0, 5)

	fmt.Println(scores, len(scores), cap(scores))
	scores = append(scores, 100)
	fmt.Println("after appending 100", scores)

	scores = append(scores, 200, 3000)

	fmt.Println("after appending 200, 3000", scores)

	scores = append(scores, 45, 55)
	fmt.Println("after appending 45, 55", scores)
	// if in case we are exceeding capacity, go grows the backing array/slice grows (usually doubles)

	scores = append(scores, 60)

	fmt.Println("after appending 60", scores)

	todos := []string{"do yourself", "workout everyday"}

	more := []string{"learn golang"}

	//...
	todos = append(todos, more...)
	fmt.Println(todos)
}
