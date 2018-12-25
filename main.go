package main

import (
  "fmt"
  "net/http"
)

func main()  {
  fmt.Println("Hello, Go web Development")
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hello, Go web development")
  })
  fmt.Println(http.ListenAndServe(":5000", nil))
}