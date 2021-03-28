package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)


type task struct {
	ID int `json:"ID"`
	Name string `json:"Name"`
	Content string `json:"Content"`
}

type allTasks [] task

var tasks = allTasks {
	{
		ID:1,
		Name: "Task One",
		Content: "Some Content",
	},
}


func indexRouter(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w,"Bienvenido a mi api")

}


func main(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRouter)
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")


	log.Fatal(	http.ListenAndServe(":4000", router))
}

func getTasks(w http.ResponseWriter, r*http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask task
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Task Data")
	}

	json.Unmarshal(reqBody, &newTask)
	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)

}

func getTask (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w,"Invalid ID")
		return
	}
	for _, task := range tasks{
		if task.ID == taskID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
		}
	}
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid User ID")
		return
	}

	for i, t := range tasks {
		if t.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Fprintf(w, "The task with ID %v has been remove successfully", taskID)
		}
	}
}

func updateTask(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	var updateTask task

	if err != nil {
		fmt.Fprintf(w, "Invalid User ID")
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please Enter valid go")
		return
	}
	json.Unmarshal(reqBody, &updateTask)

	for i, t := range tasks {
		if t.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			updateTask.ID = taskID
			tasks = append(tasks, updateTask)

			fmt.Fprintf(w, "The task with ID %v has been update successfully", taskID)
		}
	}

}
