package main
import "fmt"

func main() {

	var nome string = "ADMS Informática"
	var ano int = 1979
	sobre := fmt.Sprintf("Empresa %s Fundada em %d", nome, ano)
	fmt.Println(sobre)
}

//Concatenação usando a Função Sprintf
