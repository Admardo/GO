package main
import "fmt"

func main() {
	soma := func(x int, y int) int {
		return x + y
	} (2,2)
	fmt.Println(soma)
}

//As chamadas Funções Anônimas são aquelas que não estão vinculadas a um nome específico, que seria usado como referência para chamar essa função.
