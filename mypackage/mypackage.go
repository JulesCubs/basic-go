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
