package main

import (
	"fmt"
	"github.com/panutat-p/line-remoteinterview-go/health"
	"log"
	"net/http"
)

const PORT = 8080

func main() {
	http.HandleFunc("/", health.CheckHealth)
	fmt.Println("starting handler GET /")
	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("🟧 logging")
}
