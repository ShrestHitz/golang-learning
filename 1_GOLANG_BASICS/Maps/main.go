//MAPS
// A map is a reference type that stores key-value pairs

package main

import "fmt"

func main() {
	//creation
	//first method
	//m := make(map[string]int) //[key data type] value data type, memory already allocated
	//m["age"] = 22
	//fmt.Println(m)

	//second method

	// m := map[string]string{
	// 	"name":   "Shresth",
	// 	"salary": "2380000",
	// }
	// fmt.Println(m)

	//nil map
	//var m map[string]int
	//m["age"] = 10
	// 	m := make(map[string]int)
	// 	m["age"] = 22
	// 	m["salary"] = 22
	// 	m["num"] = 22
	// 	m["street_num"] = 22
	// 	//fmt.Println(m["nsj"])

	// 	//map of structs
	// 	type User struct {
	// 		Name string
	// 		Age  int
	// 	}
	// 	users := make(map[int]User)
	// 	users[1] = User{"Shresth", 22}
	// 	fmt.Println(users)
	// }

	points := map[string]int{
		"a": 10,
		"b": 0, //valid value
	}
	fmt.Println("a ", points["a"])
	fmt.Println("b ", points["b"])
	fmt.Println("c ", points["c"])

	valB, okB := points["b"]
	fmt.Println(valB, okB)

	valC, okC := points["c"]
	fmt.Println(valC, okC)

	if val, ok := points["c"]; ok {
		fmt.Println(val)
	} else {
		fmt.Println("c key is not present in the map")
	}

	prices := map[string]int{
		"xyz": 500,
		"def": 1800,
	}

	total := 0
	for item, price := range prices {
		fmt.Println(item, price)
		total = total + price
	}

	fmt.Println(total)
}
