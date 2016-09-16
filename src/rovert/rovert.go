package main

import (
 "html/template"
  "net/http"
  "image"
  "image/jpeg"
)

var (
  images = make(map[string]image.Image)
  templates *template.Template
)

func init() {
  templates = template.Must(
    template.ParseFiles("/Users/saileshdev/Desktop/rovert/static/root.html", "/Users/saileshdev/Desktop/rovert/static/editor.html"),
  )
}


func HandleRoot(w http.ResponseWriter, r* http.Request) {
    templates.ExecuteTemplate(w, "root.html", nil)
}


func HandleUpload(w http.ResponseWriter, r* http.Request) {
  file, header, _ := r.FormFile("image")
  image, _, _ := image.Decode(file)
  images[header.Filename] = image
  http.Redirect(w, r, "/editor?name=" + header.Filename, 303)
}

func HandleEditor(w http.ResponseWriter, r* http.Request) {
  templates.ExecuteTemplate(w, "editor.html", r.FormValue("name"))
}


func main() {
  http.HandleFunc("/", HandleRoot)
  http.HandleFunc("/upload", HandleUpload)
}
