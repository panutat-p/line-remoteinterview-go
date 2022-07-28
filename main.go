package main

import (
	"fmt"
	"github.com/panutat-p/line-remoteinterview-go/health"
	"log"
	"net/http"
	"time"
)

const PORT = 8080

func main() {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", PORT),
		Handler:      http.HandlerFunc(health.CheckHealth),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("🟧 handler GET /")
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
