package main

import (
  "database/sql"
  "flag"
	"log"
	"net/http"
  "os"

  _ "github.com/go-sql-driver/mysql"

"github.com/rmftelier/Asulside/pkg/models/mysql"

)

type application struct{
     errorLog *log.Logger
     infoLog *log.Logger
     blogs *mysql.BlogModel //n sei se tá certo o Blog
}


//Conexão ao servidor
func main() {
  
  //nome da flag, valor padrão e descrição
   addr := flag.String("addr", ":4000", "Porta da Rede")

   dsn := flag.String("dsn", "WbMxWD6L7R:7n9MPnl8l4@tcp(remotemysql.com)/WbMxWD6L7R?parseTime=true", "MySql DSN")
 
   flag.Parse()

   infoLog := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
  errorLog := log.New(os.Stderr, "ERRO:\t", log.Ldate|log.Ltime|log.Lshortfile)

    //adicionando dia 04/06 -> banco de dados
   db, err := openDB(*dsn)
   if err != nil{
       errorLog.Fatal(err)
   }
   defer db.Close()

    //applicação
    app := &application{
        errorLog: errorLog,
        infoLog: infoLog,
        blogs: &mysql.BlogModel{DB:db},
    }

    //servidor
    srv := &http.Server{
        Addr: *addr,
        ErrorLog: errorLog, 
        Handler: app.routes(),
    }


  infoLog.Printf("Inicializando o servidor na porta %s\n", *addr)
  
  err = srv.ListenAndServe() 
  errorLog.Fatal(err)
}


//Conectando no banco de dados 
func openDB(dsn string)(*sql.DB, error){
    db, err := sql.Open("mysql", dsn)
    if err!= nil{
        return nil, err
    }
    if err = db.Ping(); err != nil{
       return nil, err
    }
    return db, nil
}