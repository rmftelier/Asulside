package main
//Adicionei o fmt 
import (
	"net/http"
	"strconv"
  "html/template"
 // "fmt"
 // "os"
  //"time"
  //"path/filepath"
  //"io"

  "github.com/rmftelier/Asulside/pkg/models"
)

const MAX_UPLOAD_SIZE = 1024 * 1024 //mudei aqui

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

  
  //Testando a partir daqui -> 04/06 kkakakak
  	if r.Method == "GET" {
		tmpl, _ := template.ParseFiles("./ui/html/editor.html")

		tmpl.Execute(rw, "")
		return
	}

  /*
	if r.Method != "POST" {
		http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
  

  
	// 32 MB is the default used by FormFile
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	// get a reference to the fileHeaders
	files := r.MultipartForm.File["file"]

	for _, fileHeader := range files {
		if fileHeader.Size > MAX_UPLOAD_SIZE {
			http.Error(rw, fmt.Sprintf("The uploaded image is too big: %s. Please use an image less than 1MB in size", fileHeader.Filename), http.StatusBadRequest)
			return
		}

		file, err := fileHeader.Open()
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		defer file.Close()

		buff := make([]byte, 512)
		_, err = file.Read(buff)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		filetype := http.DetectContentType(buff)
		if filetype != "image/jpeg" && filetype != "image/png" {
			http.Error(rw, "The provided file format is not allowed. Please upload a JPEG or PNG image", http.StatusBadRequest)
			return
		}

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

    
		err = os.MkdirAll("./ui/uploads", os.ModePerm)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		f, err := os.Create(fmt.Sprintf("./ui/uploads/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename)))
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		defer f.Close()
	}

	fmt.Fprintf(rw, "Upload successful")

 */
} 






