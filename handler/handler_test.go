package handler

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTodoSuccess(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/todo?id=1", nil)
	recorder := httptest.NewRecorder()

	GetTodoHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	checkResponseCode(t, http.StatusOK, response.StatusCode)

	if body := bodyString; body == "{}" {
		t.Errorf(`Expected an empty array. Got %s`, body)
	}
}

func TestListTodoSuccess(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/todo", nil)
	recorder := httptest.NewRecorder()

	GetTodoHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	checkResponseCode(t, http.StatusOK, response.StatusCode)

	if body := bodyString; body == "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestCreateTodoSuccess(t *testing.T) {
	var jsonStr = []byte(`{"Id":4, "Task": "Makan Siang"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/todo", bytes.NewBuffer(jsonStr))
	recorder := httptest.NewRecorder()

	CreateTodoHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	checkResponseCode(t, http.StatusOK, response.StatusCode)

	expectedResponse := `{"message": "Success to create todo"}`

	if body := bodyString; body != expectedResponse {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestUpdateTodoSuccess(t *testing.T) {
	var jsonStr = []byte(`{"Task": "Berpakaian"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/todo?id=1", bytes.NewBuffer(jsonStr))
	recorder := httptest.NewRecorder()

	UpdateTodoHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	checkResponseCode(t, http.StatusOK, response.StatusCode)

	expectedResponse := `{"message": "Success to update todo"}`

	if body := bodyString; body != expectedResponse {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestDeleteTodoSuccess(t *testing.T) {
	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/todo?id=1", nil)
	recorder := httptest.NewRecorder()

	DeleteTodoHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	checkResponseCode(t, http.StatusOK, response.StatusCode)

	expectedResponse := `{"message": "Success to delete todo"}`

	if body := bodyString; body != expectedResponse {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
