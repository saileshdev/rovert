package main

import (
 "html/template"
  "net/http"
  "image"
  "image/jpeg"
)

func HandleRoot(w http.ResponseWriter, r* http.Request) {
    templates.ExecuteTemplate(w, "root.html", nil)
}

func main() {
  http.HandleFunc("/", HandleRoot)
}
