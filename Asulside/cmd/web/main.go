package main

import (
  "flag"
	"log"
	"net/http"
  "os"
)

type application struct{
     errorLog *log.Logger
     infoLog *log.Logger
}


//Conexão ao servidor
func main() {
  
  //nome da flag, valor padrão e descrição
   addr := flag.String("addr", ":4000", "Porta da Rede")
   flag.Parse()

   infoLog := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
  errorLog := log.New(os.Stderr, "ERRO:\t", log.Ldate|log.Ltime|log.Lshortfile)


//application
app := &application{
    errorLog: errorLog,
    infoLog: infoLog,
}

  //conexão ao servidor
	mux := http.NewServeMux()

  mux.HandleFunc("/", app.home)
  mux.HandleFunc("/blog", app.showBlog)
  mux.HandleFunc("/editor", app.createBlog)

  //
  fileServer := http.FileServer(http.Dir("./ui/static/"))
  mux.Handle("/static/", http.StripPrefix("/static", fileServer))


  infoLog.Printf("Inicializando o servidor na porta %s\n", *addr)
  
  err := http.ListenAndServe(*addr, mux) 
  errorLog.Fatal(err)
}