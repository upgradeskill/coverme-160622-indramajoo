package main

import (
	"log"
	"net/http"
	"task1/handler"
)

func HandleRequest(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	rw.Header().Add("Access-Control-Allow-Origin", "*")

	switch r.Method {
	case http.MethodGet:
		handler.GetTodoHandler(rw, r)
	case http.MethodPost:
		handler.CreateTodoHandler(rw, r)
	case http.MethodPut:
		handler.UpdateTodoHandler(rw, r)
	case http.MethodDelete:
		handler.DeleteTodoHandler(rw, r)
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
		rw.Write([]byte(`{"message": "Method not allowed"}`))
	}
}

func main() {
	http.HandleFunc("/todo", HandleRequest)
	log.Fatal(http.ListenAndServe("127.0.0.1:3000", nil))
}
