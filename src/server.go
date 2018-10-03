package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/foresta/go-todo-api/src/memory"
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
	category_repository := memory.NewCategoryRepository()
	task_repository := memory.NewTaskRepository()

	initialize_data()

	r := mux.NewRouter()

	// get todo list
	r.HandleFunc("/lists", func(w http.ResponseWriter, r *http.Request) {

		lists := list_repository.FindAll()

		respondJSON(w, http.StatusOK, lists)
	}).Methods("GET")

	// create new todo list
	r.HandleFunc("/lists", func(w http.ResponseWriter, r *http.Request) {
		respondJSON(w, http.StatusCreated, "{ msg: not implementation. }")
	}).Methods("POST")

	// get tasks in list
	r.HandleFunc("/lists/{lid:[0-9]+}/tasks", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		list_id, _ := strconv.Atoi(vars["lid"])
		tasks := []*Task{}

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

func initialize_data() {

	// category
	category_repository.Store(&Category{ID: 1, Title: "Work"})
	category_repository.Store(&Category{ID: 2, Title: "Private"})
	category_repository.Store(&Category{ID: 3, Title: "Development"})
	category_repository.Store(&Category{ID: 4, Title: "Design"})
	category_repository.Store(&Category{ID: 5, Title: "Art"})

	tasks = append(tasks, &Task{ID: 1, ListID: list_id, Title: "Task1", Completed: true})
	tasks = append(tasks, &Task{ID: 2, ListID: list_id, Title: "Task2", Completed: false})
	tasks = append(tasks, &Task{ID: 3, ListID: list_id, Title: "Task3", Completed: false})
	tasks = append(tasks, &Task{ID: 4, ListID: list_id, Title: "Task4", Completed: false})
	tasks = append(tasks, &Task{ID: 5, ListID: list_id, Title: "Task5", Completed: true})

	// task
	task_repository.Store(&Task{ID: 1, ListID: 1, Title: "Work Task 1", Complete: false })
	task_repository.Store(&Task{ID: 2, ListID: 2, Title: "Private Task 1",Complete: false })
	task_repository.Store(&Task{ID: 3, ListID: 3, Title: "Development Task 1" ,Complete: false })
	task_repository.Store(&Task{ID: 4, ListID: 4, Title: "Design Task 1" ,Complete: true  })
	task_repository.Store(&Task{ID: 5, ListID: 5, Title: "Art Task 1" ,Complete: false })
	task_repository.Store(&Task{ID: 6, ListID: 1, Title: "Work Task 2" ,Complete: false })
	task_repository.Store(&Task{ID: 7, ListID: 2, Title: "Private Task 2" ,Complete: false })
	task_repository.Store(&Task{ID: 8, ListID: 3, Title: "Development Task 2" ,Complete: true  })
	task_repository.Store(&Task{ID: 9, ListID: 4, Title: "Design Task 2" ,Complete: false })
	task_repository.Store(&Task{ID: 10, ListID: 5, Title: "Art Task 2" ,Complete: false })
	task_repository.Store(&Task{ID: 11, ListID: 1, Title: "Work Task 3" ,Complete: false })
	task_repository.Store(&Task{ID: 12, ListID: 2, Title: "Private Task 3" ,Complete: true  })
	task_repository.Store(&Task{ID: 13, ListID: 3, Title: "Development Task 3" ,Complete: false })
	task_repository.Store(&Task{ID: 14, ListID: 4, Title: "Design Task 3" ,Complete: false })
	task_repository.Store(&Task{ID: 15, ListID: 5, Title: "Art Task 3" ,Complete: false })
	task_repository.Store(&Task{ID: 16, ListID: 1, Title: "Work Task 4" ,Complete: true  })
	task_repository.Store(&Task{ID: 17, ListID: 2, Title: "Private Task 4" ,Complete: false })
	task_repository.Store(&Task{ID: 18, ListID: 3, Title: "Development Task 4" ,Complete: false })
	task_repository.Store(&Task{ID: 19, ListID: 4, Title: "Design Task 4" ,Complete: false })
	task_repository.Store(&Task{ID: 20, ListID: 5, Title: "Art Task 4" ,Complete: true  })
	task_repository.Store(&Task{ID: 21, ListID: 1, Title: "Work Task 5" ,Complete: false })
	task_repository.Store(&Task{ID: 22, ListID: 2, Title: "Private Task 5" ,Complete: false })
}
