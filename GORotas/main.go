package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", PaginaInicial)
	http.HandleFunc("/sobre", Sobre)
	fmt.Println("Iniciando Servidor na Porta 8080")
	http.ListenAndServe(":8080", nil)
}

func PaginaInicial(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Página Inicial")
}

func Sobre(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sobre")
}
