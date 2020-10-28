package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dtbell99/golangexamples/localdatabase"
	"github.com/gorilla/mux"
)

type health struct {
	Status     string `json:"status"`
	SystemTime string `json:"systemTime"`
}

type successResponse struct {
	Status string `json:"status"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		healthStatus := health{Status: "ok", SystemTime: time.Now().String()}
		healthJSON, err := json.Marshal(healthStatus)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(healthJSON)
	})

	r.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		allLogs := localdatabase.FindAllLogMessage()
		body, err := json.Marshal(allLogs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}).Methods("GET")

	r.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		var lm localdatabase.LogMessage
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&lm)
		if err != nil {
			panic(err)
		}
		localdatabase.AddLogMessage(lm.Message)
		sr := successResponse{Status: "ok"}
		srJSON, err := json.Marshal(sr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(srJSON)
	}).Methods("POST")

	r.HandleFunc("/log/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		localdatabase.DeleteLogMessage(id)
		sr := successResponse{Status: "ok"}
		srJSON, err := json.Marshal(sr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(srJSON)
	}).Methods("DELETE")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	fmt.Printf("Starting server at port 3000\n")
	http.ListenAndServe(":3000", r)
}
