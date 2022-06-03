package main

import (
	"fmt"
	"net/http"
	"strconv"
  "html/template"
)

func(app *application) home(rw http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/"{
    http.NotFound(rw, r)
    return
  } 
  

  files := []string{
       "./ui/html/home.html",
       "./ui/html/editor.html",
       "./ui/html/blog.html",
  }

  ts, err := template.ParseFiles(files...)
  if err != nil{
      app.errorLog.Println(err.Error())
      http.Error(rw, "Internal Error", 500)
     return
  }
  err = ts.Execute(rw, nil)
  if err != nil{
       app.errorLog.Println(err.Error())
       http.Error(rw, "Internal Error", 500)
       return
  }
}

//http://localhost:4000/snippet?id=123
func(app *application) showBlog(rw http.ResponseWriter, r *http.Request){
  id,err := strconv.Atoi(r.URL.Query().Get("id"))
  if err != nil || id < 1 {
    http.NotFound(rw, r)
    return
  }
  fmt.Fprintf(rw, "Exibir o Blog de ID: %d", id)
}

func(app *application) createBlog(rw http.ResponseWriter, r *http.Request){
  if r.Method != "POST"{
    rw.Header().Set("Allow","POST")
    http.Error(rw, "Metodo não permitido", http.StatusMethodNotAllowed)
    return
  }

  rw.Write([]byte("Criar novo blog"))
}
