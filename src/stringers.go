package main

import (
	"fmt"
	"mypackage/mypackage.go/mypackage"
)

func main() {
	var mypc mypackage.Pc

	mypc.LoadSpecs(50, 40, "intel")

	fmt.Println(mypc)
}
