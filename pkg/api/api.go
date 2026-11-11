package api

import (
	"net/http"
	"os"
	"strings"
)

const DateForm = "20060102"

func Init() {
	http.HandleFunc("/api/task", TaskHandler)
	http.HandleFunc("/api/task/", TaskHandler)
	http.HandleFunc("/api/tasks", TaskHandler)
	http.HandleFunc("/api/tasks/", TaskHandler)
	http.HandleFunc("/api/nextdate", TaskHandler)
	http.HandleFunc("/api/nextdate/", TaskHandler)
	http.HandleFunc("/api/", TaskHandler)

	http.HandleFunc("/api/signin", signinHandler)

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
		case http.MethodDelete:
			deleteTaskHandler(w, r)
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
	case "/api/task/done":
		switch r.Method {
		case http.MethodPost:
			doneTaskHandler(w, r)
		default:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"error":"incorrect method"}`))
		}
	case "/api/nextdate":
		switch r.Method {
		case http.MethodGet:
			nextDateHandler(w, r)
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

func auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pass := os.Getenv("TODO_PASSWORD")
		if len(pass) > 0 {
			var jwtToken string
			cookie, err := r.Cookie("token")
			if err == nil {
				jwtToken = cookie.Value
			}
			expToken := generateToken(pass)
			valid := (jwtToken == expToken)
			if !valid {
				http.Error(w, "Authentication required", http.StatusUnauthorized)
				return
			}
		}
		next(w, r)
	})
}
