package main

import (
	"encoding/json"
	"net/http"
	tdl "todo/todolist"

	"github.com/gorilla/mux"
)

func getList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ToDoList)
}

func addDeal(w http.ResponseWriter, r *http.Request) {
	var deal tdl.Deal
	err := json.NewDecoder(r.Body).Decode(&deal)
	if err != nil {
		panic(err)
	}
	deal.Init()
	ToDoList.Add(&deal)
}

func createRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/get-list", getList).Methods("GET")
	r.HandleFunc("/add-deal", addDeal).Methods("PUT")
	return r
}
