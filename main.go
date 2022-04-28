package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Hello struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		log.Print("GET '/'")
		json.NewEncoder(w).Encode(Hello{Message: "Hello World!"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	log.Printf("Server starting on %s port \n", fmt.Sprintf(":%v", port))
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		log.Printf("Server not starting, error : %s \n", err.Error())
	}
}
