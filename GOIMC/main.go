package main
import (
	"fmt"
	"net/http"
)

func PaginaInicial(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "imc.html")
}

func main() {
	http.HandleFunc("/", PaginaInicial)
	fmt.Println("Iniciando Aplicação em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}	
