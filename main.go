package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	const port = ":8080"
	// Set-up a simple handler for the /events route
	http.HandleFunc("/events", SSEHandler)

	fmt.Printf("Server is running at %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("Error on server start-up: %s\n", err)
	}
}

// SSEHandler is the handler for Server-Sent Events
func SSEHandler(w http.ResponseWriter, r *http.Request) {
	// Set the headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")

	dataToSend := []string{"Hello", "World", "From", "Server", "Sent", "Events"}

	for _, data := range dataToSend {
		// Write the data to the response, it should start with "data: "
		// and end with "\n\n"
		content := fmt.Sprintf("data: %s\n\n", data)
		w.Write([]byte(content))
		// Flush makes sure data has been sent immediately,
		// otherwise it will wait for the buffer to be filled
		w.(http.Flusher).Flush()
		time.Sleep(1 * time.Second)
	}

}
