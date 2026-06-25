package main

import (
	"fmt"
)

func main() {
	isLogged := true //inferred as boolean type
	isAdmin := false
	hasSubscription := true
	// AND &&

	canOpenDashboard := isLogged && hasSubscription
	canDeletePost := isAdmin || (isLogged && hasSubscription)
	fmt.Println(canOpenDashboard, canDeletePost)

}
