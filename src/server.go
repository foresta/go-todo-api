package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/foresta/go-todo-api/src/category"
	"github.com/foresta/go-todo-api/src/memory"
	"github.com/foresta/go-todo-api/src/task"
)

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

var category_repository category.Repository
var task_repository task.Repository

func main() {

	// dependency injection
	category_repository = memory.NewCategoryRepository()
	task_repository = memory.NewTaskRepository()

	initialize_data()

	r := mux.NewRouter()

	// get todo list
	r.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {

		categories := category_repository.FindAll()

		respondJSON(w, http.StatusOK, categories)
	}).Methods("GET")

	// create new todo list
	r.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		respondJSON(w, http.StatusCreated, "{ msg: not implementation. }")
	}).Methods("POST")

	// get tasks in list
	r.HandleFunc("/categories/{cid:[0-9]+}/tasks", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		category_id, _ := strconv.Atoi(vars["cid"])
		tasks := task_repository.FindByCategoryID(category_id)

		respondJSON(w, http.StatusOK, tasks)

	}).Methods("GET")

	// create task
	r.HandleFunc("/categories/{cid:[0-9]+}/tasks", func(w http.ResponseWriter, r *http.Request) {
		respondJSON(w, http.StatusCreated, "{ msg: task created. }")

	}).Methods("POST")

	// get task
	r.HandleFunc("/categories/{cid:[0-9]+}/tasks/{tid:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		task_id, _ := strconv.Atoi(vars["tid"])
		task_title := fmt.Sprintf("Task%d", task_id)
		task := &task.Task{ID: task_id, Title: task_title, Completed: false}

		respondJSON(w, http.StatusOK, task)
	}).Methods("GET")

	// update task
	r.HandleFunc("/categories/{cid:[0-9]+}/tasks/{tid:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {

	}).Methods("PUT")

	// delete task
	r.HandleFunc("/categories/{cid:[0-9]+}/tasks/{tid:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {

	}).Methods("DELETE")

	// complete task
	r.HandleFunc("/categories/{cid:[0-9]+}/tasks/{tid:[0-9]+}/complete", func(w http.ResponseWriter, r *http.Request) {

	}).Methods("PUT")

	// uncomplete task
	r.HandleFunc("/categories/{cid:[0-9]+}/tasks/{tid:[0-9]+}/complete", func(w http.ResponseWriter, r *http.Request) {

	}).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func initialize_data() {

	// category
	category_repository.Store(&category.Category{ID: 1, Title: "Work"})
	category_repository.Store(&category.Category{ID: 2, Title: "Private"})
	category_repository.Store(&category.Category{ID: 3, Title: "Development"})
	category_repository.Store(&category.Category{ID: 4, Title: "Design"})
	category_repository.Store(&category.Category{ID: 5, Title: "Art"})

	// task
	task_repository.Store(&task.Task{ID: 1, CategoryID: 1, Title: "Work Task 1", Completed: false})
	task_repository.Store(&task.Task{ID: 2, CategoryID: 2, Title: "Private Task 1", Completed: false})
	task_repository.Store(&task.Task{ID: 3, CategoryID: 3, Title: "Development Task 1", Completed: false})
	task_repository.Store(&task.Task{ID: 4, CategoryID: 4, Title: "Design Task 1", Completed: true})
	task_repository.Store(&task.Task{ID: 5, CategoryID: 5, Title: "Art Task 1", Completed: false})
	task_repository.Store(&task.Task{ID: 6, CategoryID: 1, Title: "Work Task 2", Completed: false})
	task_repository.Store(&task.Task{ID: 7, CategoryID: 2, Title: "Private Task 2", Completed: false})
	task_repository.Store(&task.Task{ID: 8, CategoryID: 3, Title: "Development Task 2", Completed: true})
	task_repository.Store(&task.Task{ID: 9, CategoryID: 4, Title: "Design Task 2", Completed: false})
	task_repository.Store(&task.Task{ID: 10, CategoryID: 5, Title: "Art Task 2", Completed: false})
	task_repository.Store(&task.Task{ID: 11, CategoryID: 1, Title: "Work Task 3", Completed: false})
	task_repository.Store(&task.Task{ID: 12, CategoryID: 2, Title: "Private Task 3", Completed: true})
	task_repository.Store(&task.Task{ID: 13, CategoryID: 3, Title: "Development Task 3", Completed: false})
	task_repository.Store(&task.Task{ID: 14, CategoryID: 4, Title: "Design Task 3", Completed: false})
	task_repository.Store(&task.Task{ID: 15, CategoryID: 5, Title: "Art Task 3", Completed: false})
	task_repository.Store(&task.Task{ID: 16, CategoryID: 1, Title: "Work Task 4", Completed: true})
	task_repository.Store(&task.Task{ID: 17, CategoryID: 2, Title: "Private Task 4", Completed: false})
	task_repository.Store(&task.Task{ID: 18, CategoryID: 3, Title: "Development Task 4", Completed: false})
	task_repository.Store(&task.Task{ID: 19, CategoryID: 4, Title: "Design Task 4", Completed: false})
	task_repository.Store(&task.Task{ID: 20, CategoryID: 5, Title: "Art Task 4", Completed: true})
	task_repository.Store(&task.Task{ID: 21, CategoryID: 1, Title: "Work Task 5", Completed: false})
	task_repository.Store(&task.Task{ID: 22, CategoryID: 2, Title: "Private Task 5", Completed: false})
}
