package main
import (
 "database/sql"   // Interação SQL Base de Dados
 "fmt"            
 "log"            
 "net/http"       // HTTP Cliente-Servidor
 "text/template"  // Template HTML

 _"github.com/mattn/go-sqlite3" // Driver SQLite
)

type Albuns struct {
 ID    int
 Titulo string
 Artista string
}

var DB *sql.DB

func initDB() {
 var err error
 DB, err = sql.Open("sqlite3", "./app.db") // Conexão com BD
 if err != nil {
  log.Fatal(err)
 }

 //SQL para Criar Tabela
 sqlStmt := `
 CREATE TABLE IF NOT EXISTS bandas (
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  titulo TEXT,
  artista TEXT
 );`

 _, err = DB.Exec(sqlStmt)
 if err != nil {
  log.Fatalf("Error creating table: %q: %s\n", err, sqlStmt)
 }
}

// Função SELECT
func indexHandler(w http.ResponseWriter, r *http.Request) {
 rows, err := DB.Query("SELECT * FROM bandas")
 if err != nil {
  http.Error(w, err.Error(), http.StatusInternalServerError) 
  return
 }
 defer rows.Close()

 todos := []Albuns{}
 for rows.Next() {
  var todo Albuns
  if err := rows.Scan(&todo.ID, &todo.Titulo, &todo.Artista); err != nil {
   http.Error(w, err.Error(), http.StatusInternalServerError) 
   return
  }
  todos = append(todos, todo)
 }

 //Template HTML
 tmpl := template.Must(template.New("index").Parse(`
 <!DOCTYPE html>
 <html>
 <head>
  <title>Catálogo de Albuns</title>
 </head>
 <body>
  <h2>Catálogo de Albuns</h2>
  <form action="/create" method="POST">
   Album <input type="text" name="titulo" required><p>
   Artista <input type="text" name="artista" required><p>
   <button type="submit">Adicionar</button>
  </form>
  <ul>
   {{range .}}
   <li>{{.Titulo}} by {{.Artista}}<a href="/delete?id={{.ID}}"> - Remover</a></li>
   {{end}}
  </ul>
 </body>
 </html>
 `))
 tmpl.Execute(w, todos) //Renderiza o Template com a Lista
}

//Função INSERT
func createHandler(w http.ResponseWriter, r *http.Request) {
 if r.Method == "POST" {
  //Recebe os Dados do Formulário
  titulo := r.FormValue("titulo")
  artista := r.FormValue("artista")
  _, err := DB.Exec("INSERT INTO bandas(titulo, artista) VALUES(?,?)", titulo, artista)
  if err != nil {
   http.Error(w, err.Error(), http.StatusInternalServerError)
   return
  }
  //Retornar Página Inicial após INSERT
  http.Redirect(w, r, "/", http.StatusSeeOther)
 }
}

//Função DELETE
func deleteHandler(w http.ResponseWriter, r *http.Request) {
 id := r.URL.Query().Get("id") // Get the id from the URL query parameters
 _, err := DB.Exec("DELETE FROM bandas WHERE id = ?", id)
 if err != nil {
  http.Error(w, err.Error(), http.StatusInternalServerError)
  return
 }
 //Retornar Página Inicial após DELETE
 http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
 initDB()         // Inicia BD
 defer DB.Close() // Encerra a Conexão quando Aplicação é fechada

 //Rotas
 http.HandleFunc("/", indexHandler)
 http.HandleFunc("/create", createHandler)
 http.HandleFunc("/delete", deleteHandler)

 fmt.Println("Acesso na Aplicação http://localhost:8080")
 log.Fatal(http.ListenAndServe(":8080", nil))
}
