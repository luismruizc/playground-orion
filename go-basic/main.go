// main.go

package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	message := fmt.Sprintf("Hello, World! Current Time: %s", currentTime)
	fmt.Fprintf(w, message)

}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
