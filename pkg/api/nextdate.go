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
			writeJson(w, http.StatusBadRequest, ErrResp{Error: "wrong date format"})
			return
		}
	}
	nextDate, err := date.WorkWithDate(nowTime, dateFromR, repeatFromR)
	if err != nil {
		writeJson(w, http.StatusBadRequest, ErrResp{Error: "wrong date format"})
		return
	}
	w.Header().Set("Content-type", "text/plain;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(nextDate))

}
