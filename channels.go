package main

import "fmt"

func say(text string, c chan<- string) {
	c <- text
}

func message(text string, d chan string) {
	d <- text
}

func main() {
	//channels
	c := make(chan string, 1)

	fmt.Println(("Hello"))

	go say("Bye", c)

	fmt.Println(<-c)

	//len y cap
	d := make(chan string, 2)

	d <- "Mensaje1"
	d <- "Mensaje2"

	fmt.Println(len(d), cap(d))

	//range y close
	close(d)

	// d <- "Mensaje3"  No puede agregar mas mensajes porque se cerro el channel y su longitud es de 2

	for message := range d {
		fmt.Println(message)
	}

	//select

	email1 := make(chan string)
	email2 := make(chan string)

	go message("mensaje3", email1)
	go message("mensaje4", email2)

	for i := 0; i < 2; i++ {
		select {
		case msj := <-email1:
			fmt.Println("Enmail recibido email1", msj)
		case msj := <-email2:
			fmt.Println("Enmail recibido email2", msj)
		}
	}
}
