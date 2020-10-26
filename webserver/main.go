package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dtbell99/golangexamples/localdatabase"
)

type health struct {
	Status     string `json:"status"`
	SystemTime string `json:"systemTime"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	healthStatus := health{Status: "ok", SystemTime: time.Now().String()}
	healthJSON, err := json.Marshal(healthStatus)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(healthJSON)
}

func logHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var lm localdatabase.LogMessage
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&lm)
		if err != nil {
			panic(err)
		}
		localdatabase.AddLogMessage(lm.Message)
	case "GET":
		allLogs := localdatabase.FindAllLogMessage()
		body, err := json.Marshal(allLogs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
		return

	default:
		err := errors.New("Invalid Method")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	logResponse := localdatabase.LogMessage{Message: "Success"}
	logResponseJSON, err := json.Marshal(logResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(logResponseJSON)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static")) // New code
	http.Handle("/", fileServer)                        // New code
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/log", logHandler)
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
