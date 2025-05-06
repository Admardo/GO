package main
import "fmt"

func main() {
	defer fmt.Println("Mundo")
	fmt.Println("Olá!")
}
//Na prática a instrução defer faz a mensagem "Mundo" seja exibida após "Olá!"
