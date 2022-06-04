package main

import (
	"net/http"
	"strconv"
  "html/template"

  "github.com/rmftelier/Asulside/pkg/models"
)

func(app *application) home(rw http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/"{
    app.notFound(rw)
    return
  } 

  blogs, err := app.blogs.Latest()
  if err != nil{
     app.serverError(rw, err)
     return
  }

  files := []string{
       "./ui/html/home.html",
       "./ui/html/editor.html",
       "./ui/html/blog.html",
  }

  ts, err := template.ParseFiles(files...)
  if err != nil{
     app.serverError(rw, err)
     return
  }
  err = ts.Execute(rw, blogs)
  if err != nil{
      app.serverError(rw, err)
       return
  }
}

//http://localhost:4000/snippet?id=1
func(app *application) showBlog(rw http.ResponseWriter, r *http.Request){
  id,err := strconv.Atoi(r.URL.Query().Get("id"))
  if err != nil || id < 1 {
    app.notFound(rw)
    return
  }

  s, err := app.blogs.Get(id)
  if err == models.ErrNoRecord{
      app.notFound(rw)
      return
  }else if err != nil{
     app.serverError(rw, err)
     return
  }

    files := []string{
		"./ui/html/blog.html",
		"./ui/html/editor.html",
		"./ui/html/home.html",
    }

    ts, err := template.ParseFiles(files...)
    if err != nil{
       app.serverError(rw, err)
       return
    }
    err = ts.Execute(rw, s)
    if err != nil{
       app.serverError(rw, err)
      return  
    }
}

func(app *application) createBlog(rw http.ResponseWriter, r *http.Request){
  if r.Method != "POST"{
    rw.Header().Set("Allow","POST")
    app.clientError(rw, http.StatusMethodNotAllowed)
    return
  }

  rw.Write([]byte("Criar novo blog"))
}
