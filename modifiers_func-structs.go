package main

import (
	"fmt"
	// pk para nombrar un alias
	pk "mypackage/mypackage.go/mypackage"
	// Para que reconozcan la importación usar el comando go mod init <la ruta del paquete> alli se creara un archovo llamado go.mod
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Porsche"
	myCar.Year = 2020
	// myCar.tax = "Matricula"  <--Esta no es accesible porque es privada

	fmt.Println(myCar)

	// Esta no es accesible porque es privada
	// var myOtherCar pk.CarPrivate
	// myCar.Brand = "Lotus"
	// myCar.Year = 2019

	// fmt.Println(myOtherCar)

	pk.PrintMessage("Este es un mensaje de prueba")

	// pk.PrintMessagePrivate("Este es un mensaje de prueba")

}
