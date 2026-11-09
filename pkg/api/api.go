package api

import (
	"net/http"
	"strings"
)

const DateForm = "20060102"

func Init() {
	http.HandleFunc("/api/task", TaskHandler)
	http.HandleFunc("/api/task/", TaskHandler)
	http.HandleFunc("/api/tasks", TaskHandler)
	http.HandleFunc("/api/tasks/", TaskHandler)

}

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)
	path = strings.TrimSuffix(path, "/")
	switch path {
	case "/api/task":
		switch r.Method {
		case http.MethodPost:
			addTaskHandler(w, r)
		case http.MethodPut:
			updateTaskHandler(w, r)
		case http.MethodGet:
			getTaskHandler(w, r)
		default:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error":"incorrect path"}`))
		}
	case "/api/tasks":
		switch r.Method {
		case http.MethodGet:
			tasksHandler(w, r)
		default:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error":"incorrect method"}`))
		}
	default:
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error":"incorrect path"}`))
	}
}
