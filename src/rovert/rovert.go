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

func main() {
  http.HandleFunc("/", HandleRoot)
}
