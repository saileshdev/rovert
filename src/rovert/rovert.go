package main

import (
  "net/http"
)

func HandleRoot(w http.ResponseWriter, r* http.Request) {
    templates.ExecuteTemplate(w, "root.html", nil)
}

func main() {
  http.HandleFunc("/", HandleRoot)
}
