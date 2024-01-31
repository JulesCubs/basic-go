package main

import (
	"fmt"
	"mypackage/mypackage.go/mypackage"
)

func main() {
	a := 50
	b := &a

	// El símbolo "&" accede a la dirección del espacio de memoria de la variable.
	// El símbolo "*" accede al valor que contiene ese espacio de memoria, dado el nombre de una variable o una dirección especifica.
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(b)
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)

	var mypc mypackage.Pc

	mypc.LoadSpecs(10, 20, "msi")

	mypc.Ping()

	mypc.DuplicateRAM()

	fmt.Println(mypc)

	mypc.DuplicateRAM()

	fmt.Println(mypc)

}
