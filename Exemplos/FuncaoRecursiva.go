package main
import "fmt"

func fatorial(x int) int {
	if x == 0 {
	return 1
	}
	return x * fatorial(x - 1)
}

func main() {
	resultado := fatorial(5)	
	fmt.Println(resultado)
}

//As Funões Resusivas são aquelas que apresentam chamadas, diretas ou indiretas, a si mesma, dentro da própria função.
