package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type List struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func respondJSON(w http.ResponseWriter, httpStatus int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	w.Write([]byte(response))
}

func main() {

	r := mux.NewRouter()

	// get todo list
	r.HandleFunc("/lists", func(w http.ResponseWriter, r *http.Request) {
		lists := []*List{}
		lists = append(lists, &List{ID: 1, Title: "List1"})
		lists = append(lists, &List{ID: 2, Title: "List2"})
		lists = append(lists, &List{ID: 3, Title: "List3"})
		lists = append(lists, &List{ID: 4, Title: "List4"})
		lists = append(lists, &List{ID: 5, Title: "List5"})

		respondJSON(w, http.StatusOK, lists)
	}).Methods("GET")

	// create new todo list
	r.HandleFunc("/lists", func(w http.ResponseWriter, r *http.Request) {
		respondJSON(w, http.StatusCreated, "{ msg: not implementation. }")
	}).Methods("POST")

	// get tasks in list
	r.HandleFunc("/lists/{lid:[0-9]+}/tasks", func(w http.ResponseWriter, r *http.Request) {
		tasks := []*Task{}
		tasks = append(tasks, &Task{ID: 1, Title: "Task1", Completed: true})
		tasks = append(tasks, &Task{ID: 2, Title: "Task2", Completed: false})
		tasks = append(tasks, &Task{ID: 3, Title: "Task3", Completed: false})
		tasks = append(tasks, &Task{ID: 4, Title: "Task4", Completed: false})
		tasks = append(tasks, &Task{ID: 5, Title: "Task5", Completed: true})

		respondJSON(w, http.StatusOK, tasks)

	}).Methods("GET")

	// create task
	r.HandleFunc("/lists/{lid:[0-9]+}/tasks", func(w http.ResponseWriter, r *http.Request) {
		respondJSON(w, http.StatusCreated, "{ msg: task created. }")

	}).Methods("POST")

	// get task
	r.HandleFunc("/lists/{lid:[0-9]+}/tasks/{tid:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		task_id, _ := strconv.Atoi(vars["tid"])
		task_title := fmt.Sprintf("Task%d", task_id)
		task := &Task{ID: task_id, Title: task_title, Completed: false}

		respondJSON(w, http.StatusOK, task)
	}).Methods("GET")

	// update task
	r.HandleFunc("/lists/{lid:[0-9]+}/tasks/{tid:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {

	}).Methods("PUT")

	// delete task
	r.HandleFunc("/lists/{lid:[0-9]+}/tasks/{tid:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {

	}).Methods("DELETE")

	// complete task
	r.HandleFunc("/lists/{lid:[0-9]+}/tasks/{tid:[0-9]+}/complete", func(w http.ResponseWriter, r *http.Request) {

	}).Methods("PUT")

	// uncomplete task
	r.HandleFunc("/lists/{lid:[0-9]+}/tasks/{tid:[0-9]+}/complete", func(w http.ResponseWriter, r *http.Request) {

	}).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
