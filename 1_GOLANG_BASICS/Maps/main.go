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
	m := make(map[string]int)
	m["age"] = 22
	m["salary"] = 22
	m["num"] = 22
	m["street_num"] = 22
	//fmt.Println(m["nsj"])

	//map of structs
	type User struct {
		Name string
		Age  int
	}
	users := make(map[int]User)
	users[1] = User{"Shresth", 22}
	fmt.Println(users)
}
