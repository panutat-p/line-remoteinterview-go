package health

import (
	"encoding/json"
	"log"
	"net/http"
)

func CheckHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "", http.StatusNotFound)
		return
	}

	if r.URL.Path != "/" {
		http.Error(w, "", http.StatusNotFound)
		return
	}

	res := Health{
		Message: "healthy",
	}
	err := json.NewEncoder(w).Encode(&res)
	if err != nil {
		log.Println(err)
	}
	log.Println("GET / 🟩")
}
