package main

import (
  "log"
  "net/http"
  "fmt"
  "os"

  "github.com/gorilla/mux"
  "github.com/sbofirov/go-sofia/internal/diagnostics"
)

type serverConf struct {
  port string
  router http.Handler
  name string
}

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

    diagnostics := diagnostics.NewDiagnostics()

    possibleErrors := make(chan error, 2)

    servers := []serverConf{
		{
			port:   blPort,
			router: router,
			name:   "application server",
		},
		{

			port:   diagnosticsPort,
			router: diagnostics,
			name:   "diagnostics server",
		},
	}

  for _, c := range servers {
  go func(conf serverConf) {
    log.Printf("The %s is preparing to handle connections...", conf.name)
    server := &http.Server{
      Addr:    ":" + conf.port,
      Handler: conf.router,
    }
    err := server.ListenAndServe()
    if err != nil {
      possibleErrors <- err
    }
  }(c)
}

select {
case err := <-possibleErrors:
  log.Fatal(err)
}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
  log.Print("Hello handler was called")
  fmt.Fprintf(w, http.StatusText(http.StatusOK))
}
