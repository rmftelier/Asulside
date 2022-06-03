package main

import (
	"log"
	"net/http"
)

//Conexão ao servidor
func main() {
	mux := http.NewServeMux()

  mux.HandleFunc("/", home)
  mux.HandleFunc("/blog", showBlog)
  mux.HandleFunc("/editor", createBlog)

  //
  fileServer := http.FileServer(http.Dir("./ui/static/"))
  mux.Handle("/static/", http.StripPrefix("/static", fileServer))
  
  log.Println("Inicializando o servidor na porta: 4000")
  err := http.ListenAndServe(":4000", mux)
  log.Fatal(err)
}