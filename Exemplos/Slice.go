package main
import "fmt"

func main() {
	numeros := [10]int{1,2,3,4,5,6,7,8,9,10}
	var v[]int = numeros[0:5]
	fmt.Println(v)
}

/*
func main() {
	nomes := [10]string{"ADMS", "Informática"}
	n := nomes[0:1]
	fmt.Println(n)
}
*/
