package main
import "fmt"

type Pessoa struct {
    Nome string
    Idade int
}

func (p *Pessoa) AtualizarIdade(novaIdade int) {
    p.Idade = novaIdade
}

func main() {
    pessoa := Pessoa{
        Nome: "ADMS",
        Idade: 40,
    }
    pessoa.AtualizarIdade(45)
    fmt.Println(pessoa.Idade)
}

//Pointer Receivers (modificam diretamente o valor original da Struct)
