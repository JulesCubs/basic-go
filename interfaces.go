package main

import (
	"fmt"
)

type figura2D interface {
	area() float64
}

type square struct {
	base float64
}

type rectangle struct {
	base   float64
	altura float64
}

func (c square) area() float64 {
	return c.base * c.base
}

func (r rectangle) area() float64 {
	return r.base * r.altura
}

func calcular(f figura2D) {
	fmt.Println("Area: ", f.area())
}

func main() {
	mySquare := square{base: 2}
	myRectengle := rectangle{base: 2, altura: 4}

	calcular(mySquare)
	calcular(myRectengle)

	//Lista de interfaces
	myInterface := []interface{}{"Hola", 12, 4.90}
	fmt.Println(myInterface...)

	figuras := []figura2D{square{base: 2}, rectangle{base: 2, altura: 4}}
	for _, figura := range figuras {
		fmt.Println(figura)
		calcular(figura)
	}
}
