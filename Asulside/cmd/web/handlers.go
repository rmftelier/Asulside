package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(rw http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/"{
    http.NotFound(rw, r)
    return
  } 
  rw.Write([]byte("Bem vindo ao Asulside"))
}

//http://localhost:4000/snippet?id=123
func showBlog(rw http.ResponseWriter, r *http.Request){
  id,err := strconv.Atoi(r.URL.Query().Get("id"))
  if err != nil || id < 1 {
    http.NotFound(rw, r)
    return
  }
  fmt.Fprintf(rw, "Exibir o Blog de ID: %d", id)
}

func createBlog(rw http.ResponseWriter, r *http.Request){
  if r.Method != "POST"{
    rw.Header().Set("Allow","POST")
    http.Error(rw, "Metodo não permitido", http.StatusMethodNotAllowed)
    return
  }

  rw.Write([]byte("Criar novo blog"))
}
