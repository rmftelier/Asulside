package main
//Adicionei o fmt 
import (
	"net/http"
	"strconv"
  "html/template"
  "fmt"

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

//Rota exibindo o formulário - TESTE
func(app *application) editor(rw http.ResponseWriter, r *http.Request){
  
    	if r.Method == "GET" {
		tmpl, _ := template.ParseFiles("./ui/html/editor.html")

		tmpl.Execute(rw, "")
		return
	}
}

//Submetendo o post
func(app *application) create(rw http.ResponseWriter, r *http.Request){ 
  title := r.FormValue("title")
	article := r.FormValue("article")

	id, err := app.blogs.Insert(title, article)
	if err != nil {
		app.serverError(rw, err)
		return
	}

	http.Redirect(rw, r, fmt.Sprintf("/blog?id=%d", id), http.StatusSeeOther)
  
} 

// Rota para About
func(app *application) about(rw http.ResponseWriter, r *http.Request){
  
    	if r.Method == "GET" {
		tmpl, _ := template.ParseFiles("./ui/html/about.html")

		tmpl.Execute(rw, "")
		return
	}
}
