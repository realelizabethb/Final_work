package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"
	"todo-list/date"
	"todo-list/pkg/db"
)

const DateForm = "20060102"

type ErrorResponse struct {
	Error string `json:"error"`
}
type TaskResponse struct {
	ID string `json:"id"`
}

func writeJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("err encoding %v", err)
	}
}
func checkData(task *db.Task) error {
	now := time.Now()
	today, _ := time.Parse(DateForm, now.Format(DateForm))
	todayStr := today.Format(DateForm)

	if task.Date == "" {
		task.Date = todayStr
		return nil
	}

	parsedDate, err := time.Parse(DateForm, task.Date)
	if err != nil {
		return errors.New("incorrect date formats")
	}
	if parsedDate.Before(today) {
		if task.Repeat == "" {
			task.Date = todayStr
		} else {
			currentDate := task.Date
			for {
				next, err := date.WorkWithDate(today, currentDate, task.Repeat)
				if err != nil {
					return errors.New("incorrect repeat")
				}
				currentDate = next
				nextTime, err := time.Parse(DateForm, currentDate)
				if err != nil {
					return errors.New("incorrect date formats")
				}
				if !nextTime.Before(today) {
					task.Date = currentDate
					break
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
