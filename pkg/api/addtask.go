package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"
	"todo-list/date"
	"todo-list/pkg/db"
)

type ErrorResponse struct {
	Error string `json:"error"`
}
type TaskResponse struct {
	ID string `json:"id"`
}

func writeJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
func checkData(task *db.Task) error {
	now := time.Now()
	today, _ := time.Parse("20060102", now.Format("20060102"))

	if task.Date == "" {
		task.Date = today.Format("20060102")
	}

	okey, err := time.Parse("20060102", task.Date)
	if err != nil {
		return errors.New("incorrect date formats")
	}

	var next string
	if len(task.Repeat) > 0 {
		/*if task.Repeat == "w" || task.Repeat == "m" {
			return errors.New("incorrect repeat")
		}*/
		next, err = date.WorkWithDate(today, task.Date, task.Repeat)
		if err != nil {
			return errors.New("incorrect repeat")
		}
	}

	taskDateStr := task.Date

	if !date.TrueDate(okey, today) {
		if len(task.Repeat) == 0 {
			task.Date = today.Format("20060102")
		} else {
			task.Date = next
		}
	} else if len(task.Repeat) > 0 {
		if taskDateStr != today.Format("20060102") {
			nextFromToday, err := date.WorkWithDate(today, today.Format("20060102"), task.Repeat)
			if err == nil {
				if nextFromToday < taskDateStr {
					task.Date = nextFromToday
				}
			}
		}
	}

	return nil
}
func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&task); err != nil {
		writeJson(w, http.StatusBadRequest, ErrorResponse{Error: "err with json"})
		return
	}
	if task.Title == "" {
		writeJson(w, http.StatusBadRequest, ErrorResponse{Error: "err title format"})
		return
	}
	if err := checkData(&task); err != nil {
		writeJson(w, http.StatusBadRequest, ErrorResponse{Error: "err date format"})
		return
	}
	dbID, err := db.AddTask(&task)
	if err != nil {
		writeJson(w, http.StatusInternalServerError, ErrorResponse{Error: "err add task"})
		return
	}

	writeJson(w, http.StatusOK, TaskResponse{ID: strconv.FormatInt(dbID, 10)})

}
