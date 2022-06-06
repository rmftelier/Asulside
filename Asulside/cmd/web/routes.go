package main

import "net/http"

func (app *application) routes() *http.ServeMux{

   //Conexão ao servidor
   mux := http.NewServeMux()

  //Rotas
  mux.HandleFunc("/", app.home)
  mux.HandleFunc("/blog", app.showBlog)
  mux.HandleFunc("/editor", app.editor)
  mux.HandleFunc("/criar", app.create)
  mux.HandleFunc("/about", app.about)
 
  fileServer := http.FileServer(http.Dir("./ui/static/"))
  mux.Handle("/static/", http.StripPrefix("/static", fileServer))

  return  mux 
}

