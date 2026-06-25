package main

import (
	"fmt"
)

func main() {
	var city string
	city = "London"

	var channel = "sangam" //Inferred to string, inference

	// :=, type inference

	subscribers := 500

	subscribers += 1000

	likes, comments := 100, 30

	fmt.Println(city, channel, subscribers, likes, comments)
}
