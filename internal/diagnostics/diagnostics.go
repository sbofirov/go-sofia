package diagnostics

import (
  "fmt"
  "github.com/gorilla/mux"
  "net/http"
)

func NewDiagnostics() *mux.Router {
  router := mux.NewRouter()
  router.HandleFunc("/healtz", healtz)
  router.HandleFunc("/ready", ready)
  return router
}

func healtz(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, http.StatusText(http.StatusOK))
}

func ready(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, http.StatusText(http.StatusOK))
}
