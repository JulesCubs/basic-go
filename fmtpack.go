package main

import "fmt"

func main() {
	helloMessage := "Hello"
	worldMessage := "world"

	// fmt.Println
	fmt.Println(helloMessage, worldMessage)

	// fmt.Println
	fmt.Println(helloMessage)
	fmt.Println(worldMessage)

	// fmt.Printf
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos)

	//  Sprintf
	message := fmt.Sprintf("%s tiene mas de %d cursos", nombre, cursos)
	fmt.Println(message)

	// Tipo de datos de una variable
	fmt.Printf("helloMessage %T\n", helloMessage)
	fmt.Printf("cursos %T\n", cursos)
}