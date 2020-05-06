package main

import (
	"encoding/json"
	"log"
	"net/http"
	tdl "todo/todolist"

	"github.com/gorilla/mux"
)

func getList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ToDoList)
}

func addDeal(w http.ResponseWriter, r *http.Request) {
	var deal tdl.Deal
	err := json.NewDecoder(r.Body).Decode(&deal)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		return
	}
	deal.Init()
	ToDoList.Add(&deal)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func editDeal(w http.ResponseWriter, r *http.Request) {
	qParam := mux.Vars(r)
	id := qParam["id"]
	var params map[string]string
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		return
	}
	ToDoList.Edit(id, params)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func completeDeal(w http.ResponseWriter, r *http.Request) {
	qParam := mux.Vars(r)
	id := qParam["id"]
	ToDoList.Complete(id)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func removeDeal(w http.ResponseWriter, r *http.Request) {
	qParam := mux.Vars(r)
	id := qParam["id"]
	ToDoList.Remove(id)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func findDeal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	qParam := mux.Vars(r)
	id := qParam["id"]
	deal := ToDoList.Find(id)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(*deal)
}

func filterDeal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params map[string]string
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Wrong request!"))
		return
	}
	filtered := ToDoList.Filter(params)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(filtered)
}

func sortDeal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	qParam := mux.Vars(r)
	sortBy := qParam["sortBy"]
	sortType := qParam["sortType"]
	sortedDeals := ToDoList.Sort(sortBy, sortType)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sortedDeals)
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
	r.HandleFunc("/sort-deal/{sortBy}/{sortType}", sortDeal).Methods("GET")
	return r
}
