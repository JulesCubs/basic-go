package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	// AND
	if valor1 == 1 && valor2 == 3 {
		fmt.Println("AND Es verdad")
	}

	// OR
	if valor1 == 1 || valor2 == 3 {
		fmt.Println("OR Es verdad")
	}

	// Convertir texto a numero
	value, err := strconv.Atoi("aa")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value", value)
}
