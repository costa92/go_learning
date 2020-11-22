package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(w,"Hello world!")
	})

	http.HandleFunc("/time/", func(w http.ResponseWriter, request *http.Request) {
		t := time.Now()
		timeStr := fmt.Sprintf("{\"time\":\"%s\"}",t)
		w.Write([]byte(timeStr))
	})

	http.ListenAndServe(":8080",nil)
}