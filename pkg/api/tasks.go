package api

import (
	"encoding/json"
	"net/http"
	"todo-list/pkg/db"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}
type ErrResp struct {
	Error string `json:"error"`
}

func tasksHandler(w http.ResponseWriter, _ *http.Request) {
	tasks, err := db.Tasks(50)
	if err != nil {
		writeJson(w, http.StatusInternalServerError, ErrResp{Error: "err with db"})
		return
	}
	if tasks == nil {
		tasks = []*db.Task{}
	}
	writeJson(w, http.StatusOK, TasksResp{Tasks: tasks})
}
func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		writeJson(w, http.StatusBadRequest, ErrResp{Error: "incorrect id"})
		return
	}
	task, err := db.GetTask(id)
	if err != nil {
		writeJson(w, http.StatusBadRequest, ErrResp{Error: "err with get task"})
		return
	}
	writeJson(w, http.StatusOK, task)
}

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var task db.Task
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&task); err != nil {
		writeJson(w, http.StatusBadRequest, ErrResp{Error: "incorrect json format"})
		return
	}
	if task.Title == "" {
		writeJson(w, http.StatusBadRequest, ErrResp{Error: "err title format"})
		return
	}
	if err := checkData(&task); err != nil {
		writeJson(w, http.StatusBadRequest, ErrResp{Error: "err date format"})
		return
	}
	err := db.UpdateTask(&task)
	if err != nil {
		writeJson(w, http.StatusBadRequest, ErrResp{Error: "err with update"})
		return
	}
	writeJson(w, http.StatusOK, map[string]any{})
}
