package main
import "fmt"

func converte (c float32) float32 {
	return 1.8 * c + 32
}

func main () {
	resultado := converte(40)
	fmt.Println("Temperatura em Fahrenheit é =", resultado)
}
