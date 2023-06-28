package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", streamHandler)
	http.ListenAndServe(":80", nil)
}

func streamHandler(w http.ResponseWriter, r *http.Request) {
	// Set response headers
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Set up a flusher to flush the response buffer
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	for num := 0; num < 10; num++ {

		event := fmt.Sprintf("data: %d\n\n", num)

		_, err := fmt.Fprint(w, event)
		if err != nil {
			return
		}

		flusher.Flush()
		log.Println("flushing")

		if num == 10 {
			panic("we're done here")
		}

		time.Sleep(time.Second)
	}
}
