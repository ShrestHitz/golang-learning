package main

import (
	"errors"
	"fmt"
)

// schedules a function call that run when the surrounding function returns
// doing some kind of cleanup
func main() {
	// defer resp.body.Close()

	fmt.Println("Case 1: success")
	if err := doWork(true); err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("Case 2: fail early")
	if err := doWork(false); err != nil {
		fmt.Println("error:", err)
	}
}

func doWork(success bool) error {

	//resource related work
	// start message -> resource acquired
	// cleanup message -> resource released

	fmt.Println("start: resource acquired")
	// defer will guarantee this runs at the end of this func
	// its going to run in both the paths
	//  going to run  when - success returns
	// going to run when - errors return
	defer fmt.Println("cleanup: resource released")

	if !success {
		return errors.New("something went wrong. I am returning early")
	}

	fmt.Println("work: doing something imp")
	fmt.Println("work: this work is done")
	return nil
}
