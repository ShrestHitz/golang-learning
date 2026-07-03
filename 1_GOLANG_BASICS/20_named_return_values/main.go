package main

import "fmt"

func divide(a int, b int) (john int, shresth int) {
	john = a / b
	shresth = a + b

	return //named return type
}

func main() {
	q, r := divide(10, 10)
	fmt.Println("division: ", q, "Sum: ", r)
}
