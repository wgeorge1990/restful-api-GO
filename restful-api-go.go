package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type event struct {
	ID          string   `json:"ID"`
	Title       string `json:"Title`
	Description string `json:"Description"`
}

type person struct {
	ID string `json:"ID"`
	Name string `json:"Name"`
	Age string `json: "Age"`
}

type allEvents []event

var events = allEvents{
	{
		ID:          "1",
		Title:       "Introduction to Golang",
		Description: "We are building a restful api is in Go and starting to get a feel for how it works and how to build cool and performant apps.",
	},
	{
		ID:          "2",
		Title:       "Meet friends in park for frisbee",
		Description: "Me and the crew are going to Piedmont Park to play some pickup frisbee on this beautiful Saturday",
	},
}

func main() {
	fmt.Println(events)
	router := mux.NewRouter().StrictSlash(true)
	fmt.Println("server is up and running on localhost:8080")
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/event", createEvent)
	router.HandleFunc("/events/{id}", getOneEvent)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Homepage.")
	fmt.Println("welcome to the homepage")
}

func getOneEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]
	fmt.Println(eventID)
	

	for _, singleEvent := range events {
		if singleEvent.ID == eventID {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, "Please enter new title and description in order to change or create")
	}
	json.Unmarshal(reqBody, &newEvent)
	events = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}
