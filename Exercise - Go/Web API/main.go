package main

import (
	"fmt"
	"net/http"
)

func HandleHealthCheck(w http.ResponseWriter, router *http.Request) {
	fmt.Println("HandleHealthCheck has been invoked...")
	if router.Method != http.MethodGet {
		w.Write([]byte(fmt.Sprintln("The request must be a GET...")))
		return
	}
	w.Write([]byte(fmt.Sprintln("The system is health...")))
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/example", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/example has been invoked...")
		w.Write([]byte(fmt.Sprintln("This is a demo endpoint...")))
	})

	router.HandleFunc("/health", HandleHealthCheck)

	errore := http.ListenAndServe(":8089", router)
	if errore != nil {
		panic(errore)
	}
}
