//Structs

//A struct in Go is a user-defined data type that groups multiple related fields into a single unit

package main

import "fmt"

type Student struct {
	ID   int
	Name string
	Age  int
}

func main() {
	var s Student

	fmt.Println(s)
}
