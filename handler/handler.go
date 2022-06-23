package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"task1/model"
)

var todos = []model.Todo{
	{Id: 1, Task: "Mandi"},
	{Id: 2, Task: "Sarapan"},
	{Id: 3, Task: "Kerja"},
}

func GetTodoHandler(rw http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	id, _ := strconv.Atoi(query.Get("id"))

	if id > 0 {
		for _, todo := range todos {
			if todo.Id == id {
				rw.WriteHeader(http.StatusOK)
				json.NewEncoder(rw).Encode(todo)
				return
			}
		}
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte(`{"message": "Data not found"}`))
	} else {
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(todos)
	}

}

func CreateTodoHandler(rw http.ResponseWriter, r *http.Request) {
	var todo model.Todo

	count := 1
	if len(todos) > 0 {
		count = todos[len(todos)-1].Id + 1
	}

	todo.Id = count
	todo.Task = r.FormValue("Task")
	todos = append(todos, todo)
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(`{"message": "Success to create todo"}`))
}

func UpdateTodoHandler(rw http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	id, _ := strconv.Atoi(query.Get("id"))

	for index, todo := range todos {
		if todo.Id == id {
			todos[index].Id = id
			todos[index].Task = r.FormValue("Task")
			rw.WriteHeader(http.StatusOK)
			rw.Write([]byte(`{"message": "Success to update todo"}`))
			return
		}
	}
	rw.WriteHeader(http.StatusNotFound)
	rw.Write([]byte(`{"message": "Data not found"}`))
}

func DeleteTodoHandler(rw http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id, _ := strconv.Atoi(query.Get("id"))

	for index, todo := range todos {
		if todo.Id == id {
			todos = append(todos[:index], todos[index+1:]...)
			rw.WriteHeader(http.StatusOK)
			rw.Write([]byte(`{"message": "Success to delete todo"}`))
			return
		}
	}
	rw.WriteHeader(http.StatusNotFound)
	rw.Write([]byte(`{"message": "Data not found"}`))
}
