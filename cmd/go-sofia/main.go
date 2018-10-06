package main

import (
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "fmt"
  "github.com/sbofirov/go-sofia/internal/diagnostics"
)

func main() {
    log.Print("Hello world!")

    router := mux.NewRouter()
    router.HandleFunc("/", handleRequest)
    go func () {
      err := http.ListenAndServe(":8080", router)
      if err != nil {
        log.Fatal(err)
      }
    }()

    diagnostics := diagnostics.NewDiagnostics()
    err := http.ListenAndServe(":8585", diagnostics)
    if err != nil {
      log.Fatal(err)
    }


}

func handleRequest(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, http.StatusText(http.StatusOK))
}
