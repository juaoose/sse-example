package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer wg.Done()
		streamHandler(w, r)
	})

	go func() {
		http.ListenAndServe(":80", nil)
	}()

	wg.Wait()
	log.Println("Application finished")
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

	done := make(chan bool)
	go func() {
		for num := 0; num < 10; num++ {
			event := fmt.Sprintf("data: %d\n\n", num)
			_, err := fmt.Fprint(w, event)
			if err != nil {
				return
			}
			flusher.Flush()
			log.Println("flushing")
			time.Sleep(time.Second)
		}
		done <- true
	}()

	select {
	case <-done:
		log.Println("Streaming completed")
	case <-r.Context().Done():
		log.Println("Connection closed by the client")
	}

	return
}
