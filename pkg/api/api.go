package api

import (
	"net/http"
)

const DateForm = "20060102"

func Init() {
	http.HandleFunc("/api/nextdate", nextDateHandler)
}
