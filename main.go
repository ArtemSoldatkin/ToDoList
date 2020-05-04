package main

import (
	"log"
	"net/http"
	tdl "todo/todolist"
)

// ToDoList - global to do list
var ToDoList tdl.ToDoList

func main() {
	deal := tdl.Deal{Name: "Name", Description: "Desc", Date: "10.03.2014 12:11"}
	deal.Init()
	ToDoList.Add(&deal)
	r := createRoutes()
	log.Fatal(http.ListenAndServe(":8000", r))
}
