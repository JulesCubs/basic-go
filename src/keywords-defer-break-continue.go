package main

import "fmt"

func main() {
	// Defer: La instrucción con esta keyword ejecuta de ultimo antes de que la función termine
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	// Continue / break
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		// continue se utiliza segun una condicion dada nos interesa que continue como errores controlados
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		// break en el sentido opuesto de continue
		if i == 8 {
			fmt.Println("Break")
			break
		}
	}
}
