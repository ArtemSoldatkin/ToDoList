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

func editDeal(w http.ResponseWriter, r *http.Request) {
	qParam := mux.Vars(r)
	id := qParam["id"]
	var params map[string]string
	json.NewDecoder(r.Body).Decode(&params)
	ToDoList.Edit(id, params)
}

func completeDeal(w http.ResponseWriter, r *http.Request) {
	qParam := mux.Vars(r)
	id := qParam["id"]
	ToDoList.Complete(id)
}

func removeDeal(w http.ResponseWriter, r *http.Request) {
	qParam := mux.Vars(r)
	id := qParam["id"]
	ToDoList.Remove(id)
}

func findDeal(w http.ResponseWriter, r *http.Request) {
	qParam := mux.Vars(r)
	id := qParam["id"]
	deal := ToDoList.Find(id)
	json.NewEncoder(w).Encode(*deal)
}

func filterDeal(w http.ResponseWriter, r *http.Request) {
	// pass
}

func sortDeal(w http.ResponseWriter, r *http.Request) {
	// pass
}

func createRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/get-list", getList).Methods("GET")
	r.HandleFunc("/add-deal", addDeal).Methods("POST")
	r.HandleFunc("/edit-deal/{id}", editDeal).Methods("PUT")
	r.HandleFunc("/complete-deal/{id}", completeDeal).Methods("PUT")
	r.HandleFunc("/remove-deal/{id}", removeDeal).Methods("DELETE")
	r.HandleFunc("/find-deal/{id}", findDeal).Methods("GET")
	r.HandleFunc("/filter-deal", filterDeal).Methods("GET")
	r.HandleFunc("/sort-deal", sortDeal).Methods("GET")
	return r
}
