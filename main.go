package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Schedule struct {
	Date       string `jason:"date"`
	GigAt      string `jason:"at"`
	AtTownName string `jason:"town"`
	GigUrl     string `jason:"url"`
}

var schedules []Schedule

func getGigSchedules(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(schedules)
}

func main() {

	r := mux.NewRouter()

	schedules = append(schedules, Schedule{Date: "2022/06/26", GigAt: "BnA_WALL", AtTownName: "Nihonbashi", GigUrl: "https://bnawall.com/events/music/4th-mural-wall-exhibition-party/"})
	schedules = append(schedules, Schedule{Date: "2022/07/17", GigAt: "Batica", AtTownName: "Ebisu", GigUrl: "http://www.batica.jp/schedule/baddest-nice-guys/"})

	r.HandleFunc("/api/gigschedules", getGigSchedules).Methods("GET")

	log.Fatal(http.ListenAndServe(":8888", r))

}
