package api

import (
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
