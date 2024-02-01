package mypackage

import "fmt"

// Public car acceso publico, las mayusculas iniciales tambien indican que es publico
type CarPublic struct {
	Brand string
	Year  int
	tax   string
}

type carPrivate struct {
	brand string
	year  int
}

// PrintMessage
func PrintMessage(text string) {
	fmt.Println(text)
}

func printMessagePrivate(text string) {
	fmt.Println(text)
}

type Pc struct {
	ram   int
	disk  int
	brand string
}

func (mypc *Pc) LoadSpecs(ram int, disk int, brand string) {
	mypc.ram = ram
	mypc.disk = disk
	mypc.brand = brand
}

func (mypc Pc) Ping() {
	fmt.Println(mypc, "pong")
}

func (mypc *Pc) DuplicateRAM() {
	mypc.ram = mypc.ram * 2
}

// Stringers
func (mypc Pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disco Duro, la marca es %s", mypc.ram, mypc.disk, mypc.brand)
}
