package main

import (
	"log"
	"net/http"
	"todo-list/pkg/api"
	"todo-list/pkg/db"
)

func main() {
	port := "7540"
	webDir := "./web"
	http.Handle("/", http.FileServer(http.Dir(webDir)))
	log.Printf("Serever work on %s", port)

	api.Init()

	//work with database
	dbFile := "scheduler.db"
	defer db.DB.Close()
	err := db.Init(dbFile)
	if err != nil {
		log.Fatal("err with database, $v", err)
	}

	//блок работы
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("problem eith server", err)
	}

}
