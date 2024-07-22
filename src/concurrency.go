package main

import (
	"fmt"
	"sync"
	"time"
)

func say(text string, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Println(text)
}

func main() {
	// Algo mas eficiente
	var wg sync.WaitGroup

	fmt.Println("Hello")

	wg.Add(1)

	// say("Hello", &wg)
	go say("world", &wg)

	wg.Wait()

	go func(text string) {
		fmt.Println(text)
	}("Adios")

	// No es recomendable usar time, porque no es eficiente debido a que va en contra del concepto de concurrencia
	time.Sleep(time.Second * 1)

}
