package main
import "fmt"

type Pessoa struct {
	Nome string
	Idade int
	Email string
}

func main() {
    pessoa := Pessoa{
        Nome:  "ADMS",
        Idade: 45,
        Email: "admsinformatica@adms.com.br",
    }

    fmt.Println("Nome:", pessoa.Nome)
    fmt.Println("Idade:", pessoa.Idade)
    fmt.Println("Email:", pessoa.Email)
}

