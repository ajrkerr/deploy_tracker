package main

import (
  "net/http"
  "log"
  "html/template"
  "github.com/ajrkerr/deploy_tracker/deploy_tracker"
)

var templates = template.Must(template.ParseGlob("views/*"))

func deployInfoTable(responseWriter http.ResponseWriter, request *http.Request) {
  var repository = PgDeployTrackerRepository{}

  _ = templates.ExecuteTemplate(responseWriter, "deployInfoTable", repository.Top10())
}

func main() {
  port := ":8080";
  assetPath := "/assets/"
  assetDir := "./assets"

  http.Handle(assetPath, http.StripPrefix(assetPath, http.FileServer(http.Dir(assetDir))))
  http.HandleFunc("/", deployInfoTable)

  err := http.ListenAndServe(port, nil)

  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}