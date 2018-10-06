package main

import (
  "log"
  "net/http"
  "fmt"
  "os"

  "github.com/gorilla/mux"
  "github.com/sbofirov/go-sofia/internal/diagnostics"
)

func main() {
    log.Print("Starting the application...")

    blPort := os.Getenv("PORT")
    if (len(blPort) == 0) {
      log.Fatal("BL port must be set!")
    }

    diagnosticsPort := os.Getenv("DIAG_PORT")
    if (len(diagnosticsPort) == 0) {
      log.Fatal("Diagnostics port must be set!")
    }

    router := mux.NewRouter()
    router.HandleFunc("/", handleRequest)
    go func () {
      err := http.ListenAndServe(":"+blPort, router)
      if err != nil {
        log.Fatal(err)
      }
    }()

    diagnostics := diagnostics.NewDiagnostics()
    err := http.ListenAndServe(":"+diagnosticsPort, diagnostics)
    if err != nil {
      log.Fatal(err)
    }


}

func handleRequest(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, http.StatusText(http.StatusOK))
}
