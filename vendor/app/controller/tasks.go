package controller

import (
	"encoding/json"
	//"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"app/model"

	"github.com/gorilla/mux"
)

//json package provides Decoder and Encoder types to support the common operation of reading and writing streams of JSON data
func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	tasks, err := model.Tasks()
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		panic(err)
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	CheckAndLogError(w, err)

	err = r.Body.Close()
	CheckAndLogError(w, err)

	//Unmarshel the body content to map Task struct
	if err = json.Unmarshal(body, &task); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		err = json.NewEncoder(w).Encode(err)
		CheckAndLogError(w, err)
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
	}
	t, er := model.CreateTask(task)
	CheckAndLogError(w, er)

	err = json.NewEncoder(w).Encode(t)
	CheckAndLogError(w, err)
}

func Show(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["taskId"])
	CheckAndLogError(w, err)

	t, er := model.TaskByID(id)
	CheckAndLogError(w, er)
	err = json.NewEncoder(w).Encode(t)
	CheckAndLogError(w, err)
}

func Destroy(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["taskId"])
	CheckAndLogError(w, err)

	err = model.DeleteTask(id)
	CheckAndLogError(w, err)

	w.Write([]byte("Request processed successfully.\n"))
}

func CheckAndLogError(w http.ResponseWriter, err error) {
	if err != nil {
		json.NewEncoder(w).Encode(err)
		panic(err)
	}
}
