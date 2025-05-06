package main
import "fmt"

func main() {
	x := 1
	switch x {
	case 0:
	 fmt.Println("Zero")
	case 1:
	 fmt.Println("Um")
	default:
	 fmt.Println("Número Inválido")
	}
}
