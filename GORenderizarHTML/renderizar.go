package main

import (
	"fmt"
	"net/http"
	"html/template"
)

type PaginaInicial struct {
	Mensagem string
}

func HomePage (w http.ResponseWriter, r *http.Request ) {
	t, _ := template.ParseFiles("renderizado.html")
		data := PaginaInicial {
			Mensagem: "Bem-Vindo!",
		}
	t.Execute(w, data)
}

func main() {
	http.HandleFunc("/", HomePage)
	fmt.Println("Iniciando Servidor na Porta 8080")
	http.ListenAndServe(":8080", nil)
}
