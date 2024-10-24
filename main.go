package main

import (
	"log"
	"net/http"
	"seven-hunter-test/ex"
)

func main() {
	http.HandleFunc("/question1", ex.MaxPathHandler)
	http.HandleFunc("/question2", ex.CatchMeHandler)
	http.HandleFunc("/question3", ex.MeatSummaryHandler)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
