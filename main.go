package main

import (
  "fmt"
  "net/http"
  "html/template"

  "database/sql"
  _ "github.com/mattn/go-sqlite3"

  "encoding/json"

)

type Page struct {
  Name string
  DBStatus bool
}

type SearchResult struct {
  Title string
  Author string
  Year string
  ID string
}

func main()  {
  templates := template.Must(template.ParseFiles("templates/index.html"))

  db, _ := sql.Open("sqlite3", "dev.db")

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    p := Page{Name: "Adarsh"}
    if name := r.FormValue("name"); name != "" {
      p.Name = name
    }
    p.DBStatus = db.Ping() == nil

    if err := templates.ExecuteTemplate(w, "index.html", p); err!= nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  })

  http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request){
    results := []SearchResult{
      SearchResult{"Some good people","Adarsha Bhattarai","2018","5355"},
      SearchResult{" Advanture of people","Ashish oli","2013","5895"},
    }
    encoder := json.NewEncoder(w)
    if err := encoder.Encode(results); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  })

  fmt.Println(http.ListenAndServe(":5000", nil))
}