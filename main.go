package main

import (
	"log"
	"net/http"
)

func main() {
	port := "7540"
	webDir := "./web"
	http.Handle("/", http.FileServer(http.Dir(webDir)))
	log.Printf("Serever work on %s", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("problem eith server", err)
	}

}
