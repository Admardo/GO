package main
import "fmt"

func main(){
	var mensagem [2]string
	mensagem[0] = "Olá!"
	mensagem[1] = "Mundo"
	fmt.Println(mensagem[0], mensagem[1])
	fmt.Println(mensagem)
}

/*
func main() {
	iniciais := [4]string{"A","D","M","S"}
	fmt.Println(iniciais)	
}
*/

/*
func main() {
	var p [10]int
	for i:=0; i < 10; i++ {
	   p[i] = i
	}
	fmt.Println(p)
}
*/
