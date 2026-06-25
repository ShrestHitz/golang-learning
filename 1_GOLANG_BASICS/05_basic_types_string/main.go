package main

import (
	"fmt"
	"strings"
)

func main() {
	firstName := "Shresth"
	LastName := "Sudhanshu"
	fullName := firstName + " " + LastName

	fmt.Println(fullName)

	fmt.Println(strings.ToUpper(fullName))
}
