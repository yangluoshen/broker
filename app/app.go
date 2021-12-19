package app

import (
	"encoding/json"
	"log"
	"net/http"
)

func NewHttpApp() *HttpApp {
	return &HttpApp{}
}

type HttpApp struct {
}

func (app *HttpApp) Run() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"stat": "UP"})
	})

	log.Println("listen on :8080")
	return http.ListenAndServe(":8080", mux)
}
