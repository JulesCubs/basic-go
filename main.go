package main

import "fmt"

func main() {
	// Declaracion de variables, constantes y zero values
	const pi float64 = 3.14
	const pi2 = 3.15
	fmt.Println("pi: ", pi)
	fmt.Println("pi2: ", pi2)

	// Declaración de variables enteras
	// Al declarar asi := se cre y asigna valor
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("areaCuadrado:", areaCuadrado)

	x := 10
	y := 50

	// Operaciones aritmeticos

	//Suma
	result := x + y
	fmt.Println("Suma:", result)

	// Resta
	result = y - x
	fmt.Println("Resta:", result)

	// Multiplicación
	result = y * x
	fmt.Println("Multiplación:", result)

	// Resta
	result = y / x
	fmt.Println("División:", result)

	// Modulo
	result = y % x
	fmt.Println("Modulo:", result)

	// Incremental
	x++
	fmt.Println("Incremental:", x)

	// Decremental
	x--
	fmt.Println("Decremental:", x)
}
