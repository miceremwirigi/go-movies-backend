package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

type Movies struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Runtime int `json:"runtime"`
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "Server Port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment (development/production)")
	flag.Parse()

	fmt.Println("... Running")

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		currentStatus := AppStatus{
			Status:      "Available",
			Environment: cfg.env,
			Version:     version,
		}

		jsonResponse, err := json.MarshalIndent(currentStatus, "", "\t")
		if err != nil {
			log.Println(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	})

	http.HandleFunc("/movies", func(w http.ResponseWriter, r *http.Request) {
		movies := []Movies{
			{
				ID : "1",
				Title: "Sparta",
				Runtime: 140,
			},
			{
				ID : "2",
				Title: "Dora",
				Runtime: 120,
			},
			{
				ID : "3",
				Title: "Guardians",
				Runtime: 104,
			},
		}

		jsonResponse, err := json.MarshalIndent(movies, "", "\t")
		if err != nil {
			log.Println(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.port), nil)
	if err != nil {
		log.Println(nil)
	}
}
