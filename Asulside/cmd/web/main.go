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


    //applicação
    app := &application{
        errorLog: errorLog,
        infoLog: infoLog,
    }

    //servidor
    srv := &http.Server{
        Addr: *addr,
        ErrorLog: errorLog, 
        Handler: app.routes(),
    }


  infoLog.Printf("Inicializando o servidor na porta %s\n", *addr)
  
  err := srv.ListenAndServe() 
  errorLog.Fatal(err)
}