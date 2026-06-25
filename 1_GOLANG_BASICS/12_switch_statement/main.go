package switchstatement

import (
	"fmt"
)

func main() {
	day := 3

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tue")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Unknown day")
	}
}
