package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path"
	"strconv"
)

type personInfo struct {
	ID              int    `json:"id"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	DOB             string `json:"dob"`
	AddressAndPhone string `json:"address_and_phone"`
}

var db = make(map[int]personInfo)

func createHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	var person personInfo
	json.Unmarshal(body, &person)

	db[person.ID] = person

	fmt.Println(db)
}

func getByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(path.Base(r.URL.Path))
	person := db[id]
	JsonMessage, _ := json.Marshal(person)
	w.Write(JsonMessage)
	fmt.Println(string(JsonMessage))
}

func updateByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(path.Base(r.URL.Path))
	requestBody, _ := io.ReadAll(r.Body)

	person := db[id]
	json.Unmarshal(requestBody, &person)

	db[id] = person
	fmt.Println(db)
}

func deleteByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(path.Base(r.URL.Path))
	delete(db, id)
	fmt.Println(db)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "taras":
		createHandler(w, r)
	case "POST":
		createHandler(w, r)
	case "GET":
		getByIDHandler(w, r)
	case "PUT":
		updateByIDHandler(w, r)
	case "DELETE":
		deleteByIDHandler(w, r)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/person/", http.HandlerFunc(handleRequest))

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
