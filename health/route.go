package health

import (
	"encoding/json"
	"log"
	"net/http"
)

func CheckHealth(w http.ResponseWriter, r *http.Request) {
	res := Health{
		Message: "healthy",
	}
	err := json.NewEncoder(w).Encode(&res)
	if err != nil {
		log.Println(err)
	}
	log.Println("GET / 🟩")
}
