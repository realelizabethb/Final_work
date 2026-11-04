package api

import (
	"net/http"
	"time"
	"todo-list/date"
)

func nextDateHandler(w http.ResponseWriter, r *http.Request) {
	var nowTime time.Time
	var err error

	nowFromR := r.FormValue("now")
	dateFromR := r.FormValue("date")
	repeatFromR := r.FormValue("repeat")

	if nowFromR == "" {
		nowTime = time.Now()
	} else {
		nowTime, err = time.Parse(DateForm, nowFromR)
		if err != nil {
			http.Error(w, "wrong date format", http.StatusBadRequest)
			return
		}
	}
	nextDate, err := date.WorkWithDate(nowTime, dateFromR, repeatFromR)
	if err != nil {
		http.Error(w, "err with repeat date", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "text/plain;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(nextDate))

}
