package main

import "fmt"

func isPalindromo(text string) {
	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println(("Es palindromo"))
	} else {
		fmt.Println(("No es palindromo"))
	}
}

func main() {
	//Array -> inmutables
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))

	//Slice -> mutable
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	//Slicing
	// Metodos en el slice
	// [n: ] antes de los puntos el elemento es inclusivo
	// [ :n] despues de los puntos el elemento es exclusivo
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Append -> Agrega elementos
	slice = append(slice, 7)
	fmt.Println(slice)

	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

	// Range -> Recorrer elementos
	sliceStr := []string{"Hola", "que", "mas?"}

	for _, value := range sliceStr {
		fmt.Println(value)
	}

	isPalindromo("torre")

	// Maps -> Diccionarios llave:valor
	maps := make(map[string]int)

	maps["Jose"] = 14
	maps["Pepe"] = 20

	fmt.Println(maps)

	// Recorrer un map
	for i, v := range maps {
		fmt.Println(i, v)
	}

	// encontrar valor
	value, ok := maps["Josefo"]
	fmt.Println(value, ok)

}
