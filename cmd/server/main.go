package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Daffc/Go_Chat_API/configs"
	"github.com/gin-gonic/gin"
)

func router () http.Handler {
  
  r := gin.New()
  r.Use(gin.Recovery())
  r.GET("/", func(c *gin.Context) {
    c.JSON(
      http.StatusOK,
      gin.H{
        "code": http.StatusOK,
        "message": "Bienvenido",
      },
    )
  })

  return r;
}

func main() {

  configs, err := configs.LoadEnv()
  if err != nil {
    log.Fatal(err);
  }

  log.Print("Starting Server...")
  server := &http.Server{
    Addr: configs.Port,
    Handler: router(),
    ReadTimeout: 10 * time.Second,
    WriteTimeout: 10 * time.Second,
    MaxHeaderBytes: 1 << 20,
  }
  log.Fatal(server.ListenAndServe())
}
