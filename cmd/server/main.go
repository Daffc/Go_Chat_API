package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Daffc/Go_Chat_API/config"
)

func main() {

  config, err := config.LoadEnv()
  if err != nil {
    log.Fatal(err);
  }

  mux := http.NewServeMux()

  mux.HandleFunc("GET /{$}", func (w http.ResponseWriter, r *http.Request){

    for name, headers := range r.Header {
      for _, h := range headers {
        fmt.Fprintf(w, "%v, %v\n", name, h)
      }
    }
  })

  log.Print("Starting Server...")
  server := &http.Server{
    Addr: config.Port,
    Handler: mux,
    ReadTimeout: 10 * time.Second,
    WriteTimeout: 10 * time.Second,
    MaxHeaderBytes: 1 << 20,
  }
  log.Fatal(server.ListenAndServe())
}
